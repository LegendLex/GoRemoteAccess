package misisapi

import "fmt"

type misisGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

type schLesson struct {
	SubjectID   int    `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	Teachers    []struct {
		ID   int    `json:"id"`
		Post string `json:"post"`
		Name string `json:"name"`
	} `json:"teachers"`
	Groups []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
	Rooms []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"rooms"`
	RoomID    int    `json:"room_id"`
	RoomName  string `json:"room_name"`
	Type      string `json:"type"`
	InCeiling bool   `json:"in_ceiling"`
	Other     *int   `json:"other"`
	Other1    *int   `json:"other1"`
}

type schDay struct {
	Type    string      `json:"type"`
	Lessons []schLesson `json:"lessons"`
}

type schHeader struct {
	Type        string `json:"type"`
	StartLesson string `json:"start_lesson"`
	EndLesson   string `json:"end_lesson"`
}

type schBell struct {
	Day1   schDay    `json:"day_1"`
	Day2   schDay    `json:"day_2"`
	Day3   schDay    `json:"day_3"`
	Day4   schDay    `json:"day_4"`
	Day5   schDay    `json:"day_5"`
	Day6   schDay    `json:"day_6"`
	Header schHeader `json:"header"`
}

type scheduleHeaderDay struct {
	Type      string `json:"type"`
	Text      string `json:"text"`
	ShortText string `json:"short_text"`
	Date      string `json:"date"`
}

type misisSchedule struct {
	Status         string `json:"status"`
	StartDate      string `json:"start_date"`
	StartDateShow  string `json:"start_date_show"`
	EndDate        string `json:"end_date"`
	EndDateShow    string `json:"end_date_show"`
	PrevDate       string `json:"prev_date"`
	NextDate       string `json:"next_date"`
	TeacherID      *int   `json:"teacher_id"`
	GroupID        string `json:"group_id"`
	RoomID         *int   `json:"room_id"`
	ScheduleHeader struct {
		Header struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"header"`
		Day1 scheduleHeaderDay `json:"day_1"`
		Day2 scheduleHeaderDay `json:"day_2"`
		Day3 scheduleHeaderDay `json:"day_3"`
		Day4 scheduleHeaderDay `json:"day_4"`
		Day5 scheduleHeaderDay `json:"day_5"`
		Day6 scheduleHeaderDay `json:"day_6"`
	} `json:"schedule_header"`
	Schedule struct {
		Bell1 schBell `json:"bell_1"`
		Bell2 schBell `json:"bell_2"`
		Bell3 schBell `json:"bell_3"`
		Bell4 schBell `json:"bell_4"`
		Bell5 schBell `json:"bell_5"`
		Bell6 schBell `json:"bell_6"`
	} `json:"schedule"`
}

// Returns string representation of the schedule for weekday
func (schedule misisSchedule) GetDay(weekday int) string {
	if weekday < 1 || weekday > 7 {
		panic("Weekday must be in [1,7]")
	}
	var (
		dayHeader scheduleHeaderDay
		lessons   []schLesson = make([]schLesson, 6)
	)
	dayHeader = scheduleHeaderDay{Text: "Воскресенье", Date: "выходной"}
	switch weekday {
	case 1:
		dayHeader = schedule.ScheduleHeader.Day1
		schedule.Schedule.Bell1.Day1.catchLessons(&lessons, 1)
		schedule.Schedule.Bell2.Day1.catchLessons(&lessons, 2)
		schedule.Schedule.Bell3.Day1.catchLessons(&lessons, 3)
		schedule.Schedule.Bell4.Day1.catchLessons(&lessons, 4)
		schedule.Schedule.Bell5.Day1.catchLessons(&lessons, 5)
		schedule.Schedule.Bell6.Day1.catchLessons(&lessons, 6)
	case 2:
		dayHeader = schedule.ScheduleHeader.Day2
		schedule.Schedule.Bell1.Day2.catchLessons(&lessons, 1)
		schedule.Schedule.Bell2.Day2.catchLessons(&lessons, 2)
		schedule.Schedule.Bell3.Day2.catchLessons(&lessons, 3)
		schedule.Schedule.Bell4.Day2.catchLessons(&lessons, 4)
		schedule.Schedule.Bell5.Day2.catchLessons(&lessons, 5)
		schedule.Schedule.Bell6.Day2.catchLessons(&lessons, 6)
	case 3:
		dayHeader = schedule.ScheduleHeader.Day3
		schedule.Schedule.Bell1.Day3.catchLessons(&lessons, 1)
		schedule.Schedule.Bell2.Day3.catchLessons(&lessons, 2)
		schedule.Schedule.Bell3.Day3.catchLessons(&lessons, 3)
		schedule.Schedule.Bell4.Day3.catchLessons(&lessons, 4)
		schedule.Schedule.Bell5.Day3.catchLessons(&lessons, 5)
		schedule.Schedule.Bell6.Day3.catchLessons(&lessons, 6)
	case 4:
		dayHeader = schedule.ScheduleHeader.Day4
		schedule.Schedule.Bell1.Day4.catchLessons(&lessons, 1)
		schedule.Schedule.Bell2.Day4.catchLessons(&lessons, 2)
		schedule.Schedule.Bell3.Day4.catchLessons(&lessons, 3)
		schedule.Schedule.Bell4.Day4.catchLessons(&lessons, 4)
		schedule.Schedule.Bell5.Day4.catchLessons(&lessons, 5)
		schedule.Schedule.Bell6.Day4.catchLessons(&lessons, 6)
	case 5:
		dayHeader = schedule.ScheduleHeader.Day5
		schedule.Schedule.Bell1.Day5.catchLessons(&lessons, 1)
		schedule.Schedule.Bell2.Day5.catchLessons(&lessons, 2)
		schedule.Schedule.Bell3.Day5.catchLessons(&lessons, 3)
		schedule.Schedule.Bell4.Day5.catchLessons(&lessons, 4)
		schedule.Schedule.Bell5.Day5.catchLessons(&lessons, 5)
		schedule.Schedule.Bell6.Day5.catchLessons(&lessons, 6)
	case 6:
		dayHeader = schedule.ScheduleHeader.Day6
		schedule.Schedule.Bell1.Day6.catchLessons(&lessons, 1)
		schedule.Schedule.Bell2.Day6.catchLessons(&lessons, 2)
		schedule.Schedule.Bell3.Day6.catchLessons(&lessons, 3)
		schedule.Schedule.Bell4.Day6.catchLessons(&lessons, 4)
		schedule.Schedule.Bell5.Day6.catchLessons(&lessons, 5)
		schedule.Schedule.Bell6.Day6.catchLessons(&lessons, 6)
		// case 7:
		// 	dayHeader = scheduleHeaderDay{Text: "Воскресенье", Date: "выходной"}
	}
	lessonsString := ""
	bells := []schHeader{
		{StartLesson: " 9:00", EndLesson: "10:35"},
		{StartLesson: "10:50", EndLesson: "12:25"},
		{StartLesson: "12:40", EndLesson: "14:15"},
		{StartLesson: "14:30", EndLesson: "16:05"},
		{StartLesson: "16:20", EndLesson: "17:55"},
		{StartLesson: "18:10", EndLesson: "19:45"},
	}
	for i, lesson := range lessons {
		if lesson.SubjectName != "" {
			lessonsString += fmt.Sprintf("\n%v| %v\n", bells[i].StartLesson, lesson.SubjectName)
		}
		if len(lesson.Teachers) > 0 {
			lessonsString += fmt.Sprintf("%v| %v\t%v\n", bells[i].EndLesson, lesson.Teachers[0].Name, lesson.RoomName)
		}
	}
	fmt.Printf("Расписание на %v: %v\n%v", dayHeader.Date, dayHeader.Text, lessonsString)
	return fmt.Sprintf("Расписание на %v: %v\n%v", dayHeader.Date, dayHeader.Text, lessonsString)
}

func (day schDay) catchLessons(lessons *[]schLesson, lessonNumber int) {
	if len(day.Lessons) > 0 {
		(*lessons)[lessonNumber-1] = day.Lessons[0]
	}
}
