# calendar-redux
trying out go

Models:
Event:
    ID: string
    Time: string
    Name: string

API:
GET api/BusyDays?year={year}&month={month}
Response: Array of Event

GET api/Events?year={year}&month={month}&day={day}
Response: Array of Event

POST api/Event?year={year}&month={month}&day={day}
Body: Event

PUT api/Event?year={year}&month={month}&day={day}
Body: Event

DELETE api/Event?year={year}&month={month}&day={day}&id={id}