package wrapper

type SongRequestWrapper struct {
	UserID    string `json:"user_id"`   // คอลัมน์ user_id ใช้เป็น foreign key
	NameSong  string `json:"name_song"` // คอลัมน์ name_song ใช้เก็บชื่อเพลง
	NameUser  string `json:"nameUser`
	URL       string `json:"url"`       // คอลัมน์ url ใช้เก็บ URL ของเพลง
	Message   string `json:"message"`   // คอลัมน์ message ใช้เก็บข้อความ
	Timestamp string `json:"timestamp"` // คอลัมน์ timestamp ใช้เก็บข้อมูลเวลา
	State     string `json:"state"`     // คอลัมน์ state ใช้เก็บสถานะ
}
