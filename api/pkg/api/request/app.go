package request

import "mime/multipart"

type UploadParams struct {
	File             *multipart.FileHeader `form:"file" binding:"required"`
	ChunkNumber      uint64                `form:"chunk_number" binding:"required,numeric"`
	ChunkSize        uint64                `form:"chunk_size" binding:"required,numeric"`
	CurrentChunkSize uint64                `form:"current_chunk_size" binding:"required,numeric"`
	TotalSize        uint64                `form:"total_size" binding:"required,numeric"`
	Identifier       string                `form:"identifier" binding:"required"`
	Filename         string                `form:"filename" binding:"required"`
	ChunkTotal       uint32                `form:"chunk_total" binding:"required,numeric"`
}

type AppInfoParams struct {
	TotalSize  uint64 `json:"total_size" binding:"required,numeric"`
	Identifier string `json:"identifier" binding:"required"`
	Filename   string `json:"filename" binding:"required"`
	ChunkTotal uint32 `json:"chunk_total" binding:"required,numeric"`
}
