package hosts

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/goodhosts/hostsfile"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	dir, err := ioutil.TempDir("", "hosts")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	hostsFile := filepath.Join(dir, "hosts")
	assert.NoError(t, ioutil.WriteFile(hostsFile, []byte(`127.0.0.1 entry1`), 0600))

	host := hosts(t, hostsFile)

	assert.NoError(t, host.Add("127.0.0.1", []string{"entry1", "entry2"}))
	assert.NoError(t, host.Add("127.0.0.2", []string{"entry3"}))

	content, err := ioutil.ReadFile(hostsFile)
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1 entry1 entry2\n127.0.0.2 entry3\n", string(content))
}

func TestRemove(t *testing.T) {
	dir, err := ioutil.TempDir("", "hosts")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	hostsFile := filepath.Join(dir, "hosts")
	assert.NoError(t, ioutil.WriteFile(hostsFile, []byte(`127.0.0.1 entry1 entry2`), 0600))

	host := hosts(t, hostsFile)

	assert.NoError(t, host.Remove([]string{"entry2"}))

	content, err := ioutil.ReadFile(hostsFile)
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1 entry1\n", string(content))
}

func TestClean(t *testing.T) {
	dir, err := ioutil.TempDir("", "hosts")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	hostsFile := filepath.Join(dir, "hosts")
	assert.NoError(t, ioutil.WriteFile(hostsFile, []byte(`127.0.0.1 entry1.suffix1 entry2.suffix2`), 0600))

	host := hosts(t, hostsFile)

	assert.NoError(t, host.Clean([]string{".suffix1"}))

	content, err := ioutil.ReadFile(hostsFile)
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1 entry2.suffix2\n", string(content))
}

func hosts(t *testing.T, hostsFile string) Hosts {
	file, err := hostsfile.NewCustomHosts(hostsFile)
	assert.NoError(t, err)
	return Hosts{
		File: &file,
	}
}