import React, {useEffect, useState} from 'react'


export default function Temperature(){
    const [weatherData, setWeatherData] = useState()
    
    // const apiUrl = (window.location.hostname === "127.0.0.1") ? 'localhost:8080' : "openweather-service"; 
    
    useEffect(() => {
        fetch(`http://minikube.info/api/health`)
            .then(res => console.log(res))
    }, []);

    return(<div>{weatherData ? weatherData : <p>No weather data available</p>}</div>);
}
