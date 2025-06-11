package entities

type Schedule struct {
	Common
	ID       int      `gorm:"primaryKey;index" json:"id" form:"id"`
	Name     string   `gorm:"type:varchar(50)" json:"name" form:"name"`
	CronTime string   `gorm:"type:varchar(50)" json:"cron_time" form:"cron_time"`
	To       To       `gorm:"serializer:json"`
	CC       CC       `gorm:"serializer:json"`
	BCC      BCC      `gorm:"serializer:json"`
	Subject  string   `gorm:"type:varchar(255)" json:"subject" form:"subject"`
	Body     string   `gorm:"type:varchar(255)" json:"body" form:"body"`
	CronID   int      `gorm:"type:int" json:"cron_id" form:"cron_id"`
	Enable   bool     `gorm:"type:bool" json:"enable" form:"enable"`
	Reports  []Report `gorm:"many2many:reports_schedules;foreignKey:ID;reference:ID;"`
	// CronList CronList
}

type To []string
type CC []string
type BCC []string

// foreignKey:InstanceID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;

type ReportsSchedules struct {
	Common
	ReportID   int `gorm:"primaryKey" form:"report_id"`
	ScheduleID int `gorm:"primaryKey" form:"schedule_id"`
}

type CronList struct {
	Common
	EntryID    int `gorm:"primaryKey" form:"entry_id"`
	ScheduleID int `gorm:"primaryKey" form:"schedule_id"`
}

// type Entry struct {
// 	// ID is the cron-assigned ID of this entry, which may be used to look up a
// 	// snapshot or remove it.
// 	ID EntryID

// 	// Schedule on which this job should be run.
// 	Schedule Schedule_cron

// 	// Next time the job will run, or the zero time if Cron has not been
// 	// started or this entry's schedule is unsatisfiable
// 	Next time.Time

// 	// Prev is the last time this job was run, or the zero time if never.
// 	Prev time.Time

// 	// Job is the thing to run when the Schedule is activated.
// 	Job Job
// }

// type FuncJob func()

// type EntryID int

// type Job interface {
// 	Run()
// }

// type Schedule_cron interface {
// 	// Next returns the next activation time, later than the given time.
// 	// Next is invoked initially, and then each time the job is run.
// 	Next(time.Time) time.Time
// }