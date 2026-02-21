package api

import milky_types "github.com/aK1r4z/emi-core/types"

type (
	// 上传私聊文件
	UploadPrivateFileRequest struct {
		UserID   int64  `json:"user_id"`   // 好友 QQ 号
		FileURI  string `json:"file_uri"`  // 文件 URI，支持 file:// http(s):// base64:// 三种格式
		FileName string `json:"file_name"` // 文件名称
	}

	UploadPrivateFileResponse struct {
		FileID string `json:"file_id"` // 文件 ID
	}

	// 上传群文件
	UploadGroupFileRequest struct {
		GroupID        int64  `json:"group_id"`         // 群号
		ParentFolderID string `json:"parent_folder_id"` // 目标文件夹 ID，默认值：/
		FileURI        string `json:"file_uri"`         // 文件 URI，支持 file:// http(s):// base64:// 三种格式
		FileName       string `json:"file_name"`        // 文件名称
	}

	UploadGroupFileResponse struct {
		FileID string `json:"file_id"` // 文件 ID
	}

	// 获取私聊文件下载链接
	GetPrivateFileDownloadURLRequest struct {
		UserID   int64  `json:"user_id"`   // 好友 QQ 号
		FileID   string `json:"file_id"`   // 文件 ID
		FileHash string `json:"file_hash"` // 文件的 TriSHA1 哈希值
	}

	GetPrivateFileDownloadURLResponse struct {
		DownloadURL string `json:"download_url"` // 文件下载链接
	}

	// 获取群文件下载链接
	GetGroupFileDownloadURLRequest struct {
		GroupID int64  `json:"group_id"` // 群号
		FileID  string `json:"file_id"`  // 文件 ID
	}

	GetGroupFileDownloadURLResponse struct {
		DownloadURL string `json:"download_url"` // 文件下载链接
	}

	// 获取群文件列表
	GetGroupFilesRequest struct {
		GroupID        int64  `json:"group_id"`         // 群号
		ParentFolderID string `json:"parent_folder_id"` // 父文件夹 ID，默认值：/
	}

	GetGroupFilesResponse struct {
		Files   []milky_types.GroupFileEntity   `json:"files"`   // 文件列表
		Folders []milky_types.GroupFolderEntity `json:"folders"` // 文件夹列表
	}

	// 移动群文件
	MoveGroupFileRequest struct {
		GroupID        int64  `json:"group_id"`         // 群号
		FileID         string `json:"file_id"`          // 文件 ID
		ParentFolderID string `json:"parent_folder_id"` // 文件所在的文件夹 ID，默认值：/
		TargetFolderID string `json:"target_folder_id"` // 目标文件夹 ID，默认值：/
	}

	MoveGroupFileResponse struct{}

	// 重命名群文件
	RenameGroupFileRequest struct {
		GroupID        int64  `json:"group_id"`         // 群号
		FileID         string `json:"file_id"`          // 文件 ID
		ParentFolderID string `json:"parent_folder_id"` // 文件所在的文件夹 ID，默认值：/
		NewFileName    string `json:"new_file_name"`    // 新文件名称
	}

	RenameGroupFileResponse struct{}

	// 删除群文件
	DeleteGroupFileRequest struct {
		GroupID int64  `json:"group_id"` // 群号
		FileID  string `json:"file_id"`  // 文件 ID
	}

	DeleteGroupFileResponse struct{}

	// 创建群文件夹
	CreateGroupFolderRequest struct {
		GroupID    int64  `json:"group_id"`    // 群号
		FolderName string `json:"folder_name"` // 文件夹名称
	}

	CreateGroupFolderResponse struct {
		FolderID string `json:"folder_id"` // 文件夹 ID
	}

	// 重命名群文件夹
	RenameGroupFolderRequest struct {
		GroupID       int64  `json:"group_id"`        // 群号
		FolderID      string `json:"folder_id"`       // 文件夹 ID
		NewFolderName string `json:"new_folder_name"` // 新文件夹名
	}

	RenameGroupFolderResponse struct{}

	// 删除群文件夹
	DeleteGroupFolderRequest struct {
		GroupID  int64  `json:"group_id"`  // 群号
		FolderID string `json:"folder_id"` // 文件夹 ID
	}

	DeleteGroupFolderResponse struct{}
)
