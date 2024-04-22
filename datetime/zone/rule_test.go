package zone

import (
	"github.com/Jumpaku/gkoku"
	"github.com/Jumpaku/gkoku/date"
	"github.com/Jumpaku/gkoku/datetime"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Transition(t *testing.T) {
	tests := []struct {
		name string
		sut  Rule
		in   int
		want Transition
	}{
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               3,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekMonday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1711963230, 0), /* 2024-04-01T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               3,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekTuesday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1712049630, 0), /* 2024-04-02T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               3,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekWednesday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1712136030, 0), /* 2024-04-03T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               3,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekThursday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1711617630, 0), /* 2024-03-28T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               3,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekFriday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1711704030, 0), /* 2024-03-29T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               3,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekSaturday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1711790430, 0), /* 2024-03-30T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               3,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekSunday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1711876830, 0), /* 2024-03-31T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               3,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekSunday,
				TimeOfDay:           datetime.TimeOf(24, 0, 0, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1711926000, 0), /* 2024-04-01T00:00:00+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               2,
				BaseDay:             1,
				DayOfWeek:           date.DayOfWeekSunday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1707038430, 0), /* 2024-02-04T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               2,
				BaseDay:             31,
				DayOfWeek:           date.DayOfWeekSunday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1709457630, 0), /* 2024-03-03T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               2,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekSunday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1709457630, 0), /* 2024-03-03T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               2,
				BaseDay:             29,
				DayOfWeek:           date.DayOfWeekSunday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2024,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1709457630, 0), /* 2024-03-03T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               2,
				BaseDay:             28,
				DayOfWeek:           date.DayOfWeekSunday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2023,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1678008030, 0), /* 2023-03-05T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
		{
			sut: NewRule(RuleArg{
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
				Month:               2,
				BaseDay:             29,
				DayOfWeek:           date.DayOfWeekSunday,
				TimeOfDay:           datetime.TimeOf(10, 20, 30, 0),
				TimeOffsetMinutes:   60,
			}),
			in: 2023,
			want: Transition{
				TransitionTimestamp: gkoku.Unix(1678008030, 0), /* 2023-03-05T10:20:30+01:00 */
				OffsetMinutesBefore: 60,
				OffsetMinutesAfter:  -60,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sut.Transition(tt.in)
			assert.Equalf(t, tt.want, got, "Transition(%v)", tt.in)
		})
	}
}
