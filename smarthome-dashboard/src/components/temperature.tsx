import React, {useEffect, useState} from 'react'


export default function Temperature(){
    const [weatherData, setWeatherData] = useState()
    useEffect(() => {
        fetch('http://localhost:8080/api/v1/weather/current')
            .then(res => res.json())
            .then(data => setWeatherData(data));
    }, []);

    return(<div>{weatherData ? weatherData : <p>No weather data available</p>}</div>);
}