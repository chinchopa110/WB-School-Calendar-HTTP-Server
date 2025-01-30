package Factory

import (
	UserServices2 "WB2/internal/Application/Contracts/UserServices"
	Parser2 "WB2/internal/Presentation/RestAPI/Parser"
	ConcreteParsers2 "WB2/internal/Presentation/RestAPI/Parser/ConcreteParsers"
)

func CreateParser(getService UserServices2.IGetService, postService UserServices2.IPostService) Parser2.IParser {
	addEventParser := &ConcreteParsers2.AddEventParser{Service: postService}
	addUserParser := &ConcreteParsers2.AddUserParser{Service: postService}
	deleteEventParser := &ConcreteParsers2.DeleteEventParser{Service: postService}
	getForDayParser := &ConcreteParsers2.GetForDayParser{Service: getService}
	getForMonthParser := &ConcreteParsers2.GetForMonthParser{Service: getService}
	getForWeekParser := &ConcreteParsers2.GetForWeekParser{Service: getService}
	updateEventDateParser := &ConcreteParsers2.UpdateEventDateParser{Service: postService}
	updateEventDescriptionParser := &ConcreteParsers2.UpdateEventDescriptionParser{Service: postService}

	baseParser := &Parser2.BaseParser{}

	return baseParser.
		AddNext(addEventParser).
		AddNext(addUserParser).
		AddNext(deleteEventParser).
		AddNext(getForDayParser).
		AddNext(getForMonthParser).
		AddNext(getForWeekParser).
		AddNext(updateEventDateParser).
		AddNext(updateEventDescriptionParser)
}
