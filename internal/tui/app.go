package tui

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/0b10headedcalf/daileet/internal/api"
	"github.com/0b10headedcalf/daileet/internal/models"
	"github.com/0b10headedcalf/daileet/internal/srs"
	"github.com/0b10headedcalf/daileet/internal/storage"
)

type screen int

const (
	screenSplash screen = iota
	screenMenu
	screenDue
	screenSolved
	screenEditor
	screenLogin
	screenReview
	screenPresets
)

type appModel struct {
	db     *sql.DB
	client *api.Client

	screen     screen
	prevScreen screen
	splash     tea.Model
	menu       tea.Model
	dueList    tea.Model
	solvedList tea.Model
	editor     tea.Model
	login      tea.Model
	review     tea.Model
	presets    tea.Model
}

// NewApp creates the root Bubble Tea model.
func NewApp(db *sql.DB) *appModel {
	client := api.NewClient()
	if session, _ := storage.GetSession(db); session != "" {
		client.SetSession(session)
	}

	return &appModel{
		db:         db,
		client:     client,
		screen:     screenSplash,
		splash:     newSplashModel(),
		menu:       newMenuModel(),
		dueList:    newProblemListModel(listDue),
		solvedList: newProblemListModel(listSolved),
		editor:     newEditorModel(),
		login:      newLoginModel(),
		presets:    newPresetsModel(),
	}
}

func (m *appModel) Init() tea.Cmd {
	return m.splash.Init()
}

func (m *appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case splashDoneMsg:
		m.screen = screenMenu
		return m, m.menu.Init()
	case menuSelectedMsg:
		switch msg.choice {
		case menuDue:
			m.screen = screenDue
			return m, m.dueList.Init()
		case menuSolved:
			m.screen = screenSolved
			return m, m.solvedList.Init()
		case menuEditor:
			m.screen = screenEditor
			return m, m.editor.Init()
		case menuLogin:
			m.screen = screenLogin
			return m, m.login.Init()
		case menuPresets:
			m.screen = screenPresets
			return m, m.presets.Init()
		}
	case goBackMsg:
		m.screen = screenMenu
		return m, nil
	case reviewProblemMsg:
		m.prevScreen = m.screen
		m.screen = screenReview
		m.review = newReviewModel(msg.problem)
		return m, m.review.Init()
	case gradeMsg:
		if err := m.applyGrade(msg.problem, msg.result); err != nil {
			// In a real app we'd show an error toast; for now just print.
			_ = err
		}
		m.screen = m.prevScreen
		return m, m.refreshCurrentScreen()
	case doRefreshListMsg:
		return m, m.loadList(msg.kind)
	case loadedProblemsMsg:
		return m.updateListModel(msg)
	case refreshEditorMsg:
		return m, m.loadEditor()
	case doRefreshEditorMsg:
		return m, m.loadEditor()
	case loadedEditorProblemsMsg:
		return m.updateEditorModel(msg.problems)
	case addProblemMsg:
		return m, m.addProblem(msg.slug)
	case addProblemSuccessMsg:
		return m.updateEditorSuccess(msg.title)
	case deleteProblemMsg:
		return m, m.deleteProblem(msg.id)
	case saveSessionMsg:
		return m, m.saveSession(msg.session)
	case loginSuccessMsg:
		// Handled by login model
	case loadPresetMsg:
		return m, m.loadPreset(msg.choice)
	case loadCustomPresetMsg:
		return m, m.loadCustomPreset(msg.path)
	case presetLoadedMsg:
		m.presets, _ = m.presets.Update(msg)
		return m, nil
	}

	// Route to active sub-model.
	switch m.screen {
	case screenSplash:
		m.splash, cmd = m.splash.Update(msg)
	case screenMenu:
		m.menu, cmd = m.menu.Update(msg)
	case screenDue:
		m.dueList, cmd = m.dueList.Update(msg)
	case screenSolved:
		m.solvedList, cmd = m.solvedList.Update(msg)
	case screenEditor:
		m.editor, cmd = m.editor.Update(msg)
	case screenLogin:
		m.login, cmd = m.login.Update(msg)
	case screenReview:
		m.review, cmd = m.review.Update(msg)
	case screenPresets:
		m.presets, cmd = m.presets.Update(msg)
	}
	return m, cmd
}

func (m *appModel) View() tea.View {
	switch m.screen {
	case screenSplash:
		return m.splash.View()
	case screenMenu:
		return m.menu.View()
	case screenDue:
		return m.dueList.View()
	case screenSolved:
		return m.solvedList.View()
	case screenEditor:
		return m.editor.View()
	case screenLogin:
		return m.login.View()
	case screenReview:
		return m.review.View()
	case screenPresets:
		return m.presets.View()
	}
	return tea.View{}
}

// --- DB-backed commands ---

func (m *appModel) refreshCurrentScreen() tea.Cmd {
	switch m.screen {
	case screenDue:
		return m.dueList.Init()
	case screenSolved:
		return m.solvedList.Init()
	case screenEditor:
		return m.editor.Init()
	}
	return nil
}

func (m *appModel) loadList(kind problemListKind) tea.Cmd {
	return func() tea.Msg {
		var probs []models.Problem
		var err error
		switch kind {
		case listDue:
			probs, err = storage.ListDue(m.db)
		case listSolved:
			probs, err = storage.ListSolved(m.db)
		}
		if err != nil {
			return loadedProblemsMsg{kind: kind, problems: nil}
		}
		return loadedProblemsMsg{kind: kind, problems: probs}
	}
}

func (m *appModel) updateListModel(msg loadedProblemsMsg) (tea.Model, tea.Cmd) {
	switch msg.kind {
	case listDue:
		m.dueList, _ = m.dueList.Update(msg)
	case listSolved:
		m.solvedList, _ = m.solvedList.Update(msg)
	}
	return m, nil
}

func (m *appModel) applyGrade(p models.Problem, result int) error {
	state := scheduler.SM2State{
		Interval:    p.Interval,
		Repetitions: p.Repetitions,
		EaseFactor:  p.EaseFactor,
	}
	nextDue := state.Apply(scheduler.ReviewResult(result))
	p.Interval = state.Interval
	p.Repetitions = state.Repetitions
	p.EaseFactor = state.EaseFactor
	p.DueDate = &nextDue
	now := time.Now()
	p.LastReviewed = &now
	if p.Repetitions >= 1 {
		p.Status = "review"
	}
	return storage.UpdateProblemReview(m.db, p)
}

func (m *appModel) loadEditor() tea.Cmd {
	return func() tea.Msg {
		probs, err := storage.ListAll(m.db)
		if err != nil {
			return loadedEditorProblemsMsg{problems: nil}
		}
		return loadedEditorProblemsMsg{problems: probs}
	}
}

func (m *appModel) updateEditorModel(probs []models.Problem) (tea.Model, tea.Cmd) {
	m.editor, _ = m.editor.Update(loadedEditorProblemsMsg{problems: probs})
	return m, nil
}

func (m *appModel) updateEditorSuccess(title string) (tea.Model, tea.Cmd) {
	m.editor, _ = m.editor.Update(addProblemSuccessMsg{title: title})
	return m, nil
}

func (m *appModel) addProblem(slug string) tea.Cmd {
	return func() tea.Msg {
		slug = strings.TrimSpace(slug)
		meta, err := m.client.FetchProblemMeta(slug)
		if err != nil {
			// Fallback: just use slug as title
			meta.Title = slug
			meta.Difficulty = "Medium"
		}
		p := models.Problem{
			Title:      meta.Title,
			TitleSlug:  slug,
			Difficulty: models.Difficulty(meta.Difficulty),
			URL:        fmt.Sprintf("https://leetcode.com/problems/%s/", slug),
			Status:     "new",
		}
		if err := storage.InsertProblem(m.db, p); err != nil {
			return addProblemSuccessMsg{title: "(error)"}
		}
		return addProblemSuccessMsg{title: p.Title}
	}
}

func (m *appModel) deleteProblem(id int) tea.Cmd {
	return func() tea.Msg {
		_ = storage.DeleteProblemByID(m.db, id)
		return refreshEditorMsg{}
	}
}

func (m *appModel) saveSession(session string) tea.Cmd {
	return func() tea.Msg {
		_ = storage.SetSession(m.db, session)
		m.client.SetSession(session)
		return loginSuccessMsg{}
	}
}

func (m *appModel) loadPreset(choice presetChoice) tea.Cmd {
	var filename string
	switch choice {
	case presetBlind75:
		filename = "blind75.json"
	case presetGrind75:
		filename = "grind75.json"
	default:
		return func() tea.Msg { return presetLoadedMsg{added: 0, failed: 0} }
	}
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, "presets", filename)
	return m.loadPresetFromPath(path)
}

func (m *appModel) loadCustomPreset(path string) tea.Cmd {
	return m.loadPresetFromPath(path)
}

func (m *appModel) loadPresetFromPath(path string) tea.Cmd {
	return func() tea.Msg {
		data, err := os.ReadFile(path)
		if err != nil {
			return presetLoadedMsg{added: 0, failed: 1}
		}
		var slugs []string
		if err := json.Unmarshal(data, &slugs); err != nil {
			// Try object array with "slug" field
			var objs []struct {
				Slug string `json:"slug"`
			}
			if err2 := json.Unmarshal(data, &objs); err2 == nil {
				for _, o := range objs {
					if o.Slug != "" {
						slugs = append(slugs, o.Slug)
					}
				}
			} else {
				return presetLoadedMsg{added: 0, failed: 1}
			}
		}
		return m.addProblemsFromSlugs(slugs)
	}
}

func (m *appModel) addProblemsFromSlugs(slugs []string) tea.Msg {
	var added, failed int
	for _, slug := range slugs {
		slug = strings.TrimSpace(slug)
		if slug == "" {
			continue
		}
		meta, err := m.client.FetchProblemMeta(slug)
		if err != nil {
			meta.Title = slug
			meta.Difficulty = "Medium"
		}
		p := models.Problem{
			Title:      meta.Title,
			TitleSlug:  slug,
			Difficulty: models.Difficulty(meta.Difficulty),
			URL:        fmt.Sprintf("https://leetcode.com/problems/%s/", slug),
			Status:     "new",
		}
		if err := storage.InsertProblem(m.db, p); err != nil {
			failed++
		} else {
			added++
		}
	}
	return presetLoadedMsg{added: added, failed: failed}
}
