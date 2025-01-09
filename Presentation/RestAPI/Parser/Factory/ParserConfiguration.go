package Factory

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Parser"
	"WB2/Presentation/RestAPI/Parser/ConcreteParsers"
)

func CreateParser(getService UserServices.IGetService, postService UserServices.IPostService) Parser.IParser {
	addEventParser := &ConcreteParsers.AddEventParser{Service: postService}
	addUserParser := &ConcreteParsers.AddUserParser{Service: postService}
	deleteEventParser := &ConcreteParsers.DeleteEventParser{Service: postService}
	getForDayParser := &ConcreteParsers.GetForDayParser{Service: getService}
	getForMonthParser := &ConcreteParsers.GetForMonthParser{Service: getService}
	getForWeekParser := &ConcreteParsers.GetForWeekParser{Service: getService}
	updateEventDateParser := &ConcreteParsers.UpdateEventDateParser{Service: postService}
	updateEventDescriptionParser := &ConcreteParsers.UpdateEventDescriptionParser{Service: postService}

	baseParser := &Parser.BaseParser{}

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
