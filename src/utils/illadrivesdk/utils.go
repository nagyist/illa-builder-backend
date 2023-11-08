package illadrivesdk

import "errors"

func ExtractRawFilesFromListResponse(listResponse map[string]interface{}) ([]map[string]interface{}, error) {
	// extract file ids
	listedFiles, hitListedFiles := listResponse["files"]
	if !hitListedFiles {
		return nil, errors.New("mossing files field in list response")
	}

	// assert sub structure
	listedFilesAsserted, assertListedFilesPass := listedFiles.([]interface{})
	if !assertListedFilesPass {
		return nil, errors.New("invalied file list returned")
	}
	files := make([]map[string]interface{}, 0)
	for _, listedFile := range listedFilesAsserted {
		listedFileAsserted, listedFileAssertePass := listedFile.(map[string]interface{})
		if !listedFileAssertePass {
			return nil, errors.New("invalied file list element returned")
		}
		files = append(files, listedFileAsserted)
	}
	return files, nil
}

func ExtractFileIDFromRawFiles(files []map[string]interface{}) ([]string, error) {
	fileIDs := make([]string, 0)
	for _, file := range files {
		fileID, hitFileID := file["id"]
		if !hitFileID {
			return nil, errors.New("missing field id from files data")
		}
		fileIDString, fileIDAssertPass := fileID.(string)
		if !fileIDAssertPass {
			return nil, errors.New("invalied file ID data type")
		}
		fileIDs = append(fileIDs, fileIDString)
	}
	return fileIDs, nil
}

func ExtendRawFilesTinyURL(files []map[string]interface{}, tinyURLsMap map[string]string) ([]map[string]interface{}, error) {
	for serial, file := range files {
		fileID, hitFileID := file["id"]
		if !hitFileID {
			return nil, errors.New("missing field id from files data for extend")
		}
		fileIDString, fileIDAssertPass := fileID.(string)
		if !fileIDAssertPass {
			return nil, errors.New("invalied file ID data type")
		}
		tinyURL, hitTinyURL := tinyURLsMap[fileIDString]
		if !hitTinyURL {
			return nil, errors.New("missing field id mapped tiny URL")

		}
		files[serial]["tinyURL"] = tinyURL

	}
	return files, nil
}