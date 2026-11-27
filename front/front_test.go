package front

import (
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Vuetify 2 VMenu measures its content only when it is opened (menuable.updateDimensions).
// If the menu was shifted up because a long project list did not fit below the "Project" button,
// it stays there after the search filter shrinks the list. The menu must be re-measured
// whenever the filtered list changes, so that it is docked to the "Project" button again.
func TestProjectsMenuRepositionsOnSearch(t *testing.T) {
	data, err := os.ReadFile("src/App.vue")
	require.NoError(t, err)
	src := string(data)

	menuTag := regexp.MustCompile(`<v-menu[^>]*v-model="projectsMenu"[^>]*>`).FindString(src)
	require.NotEmpty(t, menuTag, "projects v-menu not found")

	refMatch := regexp.MustCompile(`\sref="([^"]+)"`).FindStringSubmatch(menuTag)
	require.Len(t, refMatch, 2, "projects v-menu must have a ref to be re-measured after filtering: %s", menuTag)
	ref := regexp.QuoteMeta(refMatch[1])

	watcher := regexp.MustCompile(`(?s)\n\s+(filteredProjects|normalizedProjectSearch|projectSearch)\s*(\(\w*\)\s*)?[{(:].*?\$refs\.` + ref + `.*?\.(updateDimensions|onResize)\(`)
	assert.Regexp(t, watcher, src, "the projects menu must be re-measured (updateDimensions/onResize) when the search changes the list")
}
