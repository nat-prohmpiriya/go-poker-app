package model

type Hand struct {
	Rank     int    // ระดับมือไพ่ 1-10
	Name     string // "Flush", "One Pair", etc.
	TieBreak []int  // คำสำหรับเทีบตอนเสมอกัน
}
