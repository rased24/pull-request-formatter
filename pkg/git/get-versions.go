package git

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"pull-request-formatter/pkg/config"
	"regexp"
	"strconv"
	"strings"
)

func GetVersions() (versions []version, err error) {

	files, err := getFiles()
	if err != nil {
		return
	}

	changedFolders := findFolders(files)

	var req *http.Request
	var res *http.Response

	for _, folder := range changedFolders {
		apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s/%s/%s", config.GitOwner, config.GitRepo, config.FilePathToUpdate, folder, config.FileToUpdate)

		req, err = http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			return
		}

		res, err = send(req)
		if err != nil {
			return
		}

		var initFile fileResponse

		err = json.NewDecoder(res.Body).Decode(&initFile)
		if err != nil {
			return
		}

		err = res.Body.Close()
		if err != nil {
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(initFile.Content)
		if err != nil {
			return nil, err
		}

		re := regexp.MustCompile(config.VersionRegex)

		match := re.FindStringSubmatch(string(decoded))

		if len(match) == 0 {
			fmt.Println("malformed")
		}

		oldVersion := match[1]

		newVersion, err := getNextVersion(oldVersion)
		if err != nil {
			return nil, err
		}

		oldIntVersion, err := versionToInt(oldVersion)
		if err != nil {
			return nil, err
		}

		newIntVersion, err := versionToInt(newVersion)
		if err != nil {
			return nil, err
		}

		updateLog := version{
			Name:          folder,
			OldVersion:    oldVersion,
			NewVersion:    newVersion,
			OldIntVersion: oldIntVersion,
			NewIntVersion: newIntVersion,
		}
		versions = append(versions, updateLog)

		newBody := strings.Replace(string(decoded), oldVersion, newVersion, 1)

		commit := updatedFileBody{
			Sha:     initFile.Sha,
			Message: fmt.Sprintf("Version increase for %s from %s to %s", updateLog.Name, updateLog.OldVersion, updateLog.NewVersion),
			Committer: committer{
				Name:  config.CommitterName,
				Email: config.CommitterEmail,
			},
			Branch:  config.CommitBranch,
			Content: base64.StdEncoding.EncodeToString([]byte(newBody)),
		}

		updateContent(commit, initFile.Path)
	}

	return
}

func findFolders(files []fileBody) (changedFolders []string) {
	pattern := regexp.MustCompile(config.PatternRegex)

	for _, f := range files {
		filepath := f.Filename

		match := pattern.FindString(filepath)

		if match != "" {
			// remove the prefix
			match = strings.TrimPrefix(match, config.FilePathToUpdate+"/")

			if !inArray(changedFolders, match) {
				changedFolders = append(changedFolders, match)
			}
		}
	}
	return
}

func getFiles() (files []fileBody, err error) {
	Init()

	err = getPr()
	if err != nil {
		return
	}

	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%d/files", config.GitOwner, config.GitRepo, pr.Number)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}

	res, err := send(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	err = json.NewDecoder(res.Body).Decode(&files)

	return
}

func getNextVersion(version string) (string, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid version format")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", err
	}

	if patch < 9 {
		patch++
	} else if minor < 9 {
		minor++
		patch = 0
	} else {
		major++
		minor = 0
		patch = 0
	}

	return fmt.Sprintf("%d.%d.%d", major, minor, patch), nil
}

func versionToInt(version string) (int, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid version format: %s", version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid major version: %s", parts[0])
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minor version: %s", parts[1])
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("Invalid patch version: %s", parts[2])
	}

	return major*10000 + minor*100 + patch, nil
}

func inArray(array []string, target string) bool {
	for _, s := range array {
		if s == target {
			return true
		}
	}
	return false
}
