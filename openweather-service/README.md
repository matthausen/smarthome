## Open Weather API
This service is using the free API from OpenWeather.
There are 2 endpoints:

1. Hourly forecast fopr curremt weather
2. 5 days hourly weather forecast

How to display the icon in the frontend from the response:
From response weather[0].icon
`var iconurl = "http://openweathermap.org/img/w/" + iconcode + ".png";`