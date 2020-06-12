# calendar-redux
trying out go

TODO: db password as a secret, write actual README

Models:
Event:
    ID: string
    Time: string
    Name: string

API:
GET api/BusyDays?year={year}&month={month}
Response: Array of ints

GET api/Events?year={year}&month={month}&day={day}
Response: Array of Event

POST api/Event?year={year}&month={month}&day={day}
Body: Event

PUT api/Event?year={year}&month={month}&day={day}
Body: Event

DELETE api/Event?year={year}&month={month}&day={day}&id={id}