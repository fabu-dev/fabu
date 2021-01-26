package request

import "mime/multipart"

type AppIndexParams struct {
	TeamId uint64 `form:"team_id" binding:"required,numeric"`
}

type UploadParams struct {
	File             *multipart.FileHeader `form:"file" binding:"required"`
	ChunkNumber      int                   `form:"chunk_number" binding:"required,numeric"`
	ChunkSize        uint64                `form:"chunk_size" binding:"required,numeric"`
	CurrentChunkSize uint64                `form:"current_chunk_size" binding:"required,numeric"`
	TotalSize        uint64                `form:"total_size" binding:"required,numeric"`
	Identifier       string                `form:"identifier" binding:"required"`
	Filename         string                `form:"filename" binding:"required"`
	ChunkTotal       int                   `form:"chunk_total" binding:"required,numeric"`
}

type AppInfoParams struct {
	TotalSize  uint64 `json:"total_size" binding:"required,numeric"`
	Identifier string `json:"identifier" binding:"required"`
	Filename   string `json:"filename" binding:"required"`
	ChunkTotal uint32 `json:"chunk_total" binding:"required,numeric"`
}

type SaveParams struct {
	Identifier  string `json:"identifier" binding:"required"`
	Description string `json:"description" binding:"required"`
	TeamId      uint64 `json:"team_id" binding:"required,numeric"`
}

type AppViewParams struct {
	Id uint64 `uri:"id" binding:"required,numeric"`
}

type AppViewByShortParams struct {
	ShortUrl string `form:"short_url" binding:"required,alpha"`
}

type AppDeleteParams struct {
	Id uint64 `uri:"id" binding:"required,numeric"`
}
