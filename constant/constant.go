package constant

const (
	CreateVideosTable     = "CREATE TABLE IF NOT EXISTS videos (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, url TEXT, duration INTEGER, created_at TEXT, updated_at TEXT)"
	CreateAnnotationTable = "CREATE TABLE IF NOT EXISTS annotations (id INTEGER PRIMARY KEY AUTOINCREMENT, video_id INTEGER, start_time INTEGER, end_time INTEGER, type TEXT, notes TEXT, created_at TEXT, updated_at TEXT)"
	InsertIntoVideosTable = "INSERT INTO videos (title, url, duration, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
)
