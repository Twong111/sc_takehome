package folders

import (
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	/*
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	*/
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	/* ori
	for k, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	*/
	for _, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	for _, v1 := range f {
		fp = append(fp, &v1)
	}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()
	/* original
	resFolder := []*Folder{} //bad varible name
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	*/
	resFolders := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolders = append(resFolders, folder)
		}
	}
	return resFolders, nil
}
