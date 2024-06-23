import (
	"strings"

	"github.com/samber/lo"
)
	// Custom patches tend to work better when treating new files as diffs
	// against an empty file. The only case where we need this to be false is
	// when moving a custom patch to an earlier commit; in that case the patch
	// command would fail with the error "file does not exist in index" if we
	// treat it as a diff against an empty file.
	TurnAddedFilesIntoDiffAgainstEmptyFile bool

	} else if self.opts.TurnAddedFilesIntoDiffAgainstEmptyFile {
		result := make([]string, 0, len(self.patch.header))
		for idx, line := range self.patch.header {
			if strings.HasPrefix(line, "new file mode") {
				continue
			}
			if line == "--- /dev/null" && strings.HasPrefix(self.patch.header[idx+1], "+++ b/") {
				line = "--- a/" + self.patch.header[idx+1][6:]
			}
			result = append(result, line)
		}
		return result