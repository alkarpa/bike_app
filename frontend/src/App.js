import { useEffect, useState } from 'react';
import './App.css';
import RideTable from './components/RideTable';
import StationTable from './components/StationTable';

const API_path = "http://localhost:8080"

const getStations = async () => {
  const station_api = `${API_path}/station/`
  let response = await fetch(station_api)
  let json = await response.json()

  const stations = json //.reduce((map, cur) => ({ ...map, [cur.id]: cur.name }), {})
  return stations
}

const getRides = async (parameters = {}) => {
  const ride_api = `${API_path}/ride/`
  let response = await fetch(ride_api)
  let json = await response.json()
  return json
}

function App() {
  const [lang, setLang] = useState('fi')
  const [stations, setStations] = useState([])
  const [rides, setRides] = useState([])

  useEffect(() => {
    const fetch_data = async () => {
      let stations = await getStations()
      let initial_rides = await getRides()
      setStations(stations)
      setRides(initial_rides)
    }
    fetch_data()
  }, [])
  
  const flags = {
    "se": "🇸🇪",
    "fi": "🇫🇮",
  }

  const stationLang = stations.reduce( (map, cur) => ({ ...map, [cur.id]: cur["text"][lang] }), {} )
  return (
    <div className="App">
      <div>
        Data language:
        { ["fi","se"].map( l => (<button key={`lb_${l}`} onClick={()=>setLang(l)}>{flags[l]}</button>) )}
      </div>
      <div className="Tables">
      <div>
        <StationTable lang={lang} stations={stations} />

      </div>
      <RideTable stationLang={stationLang} rides={rides} />

      </div>
      
    </div>
  );
}

export default App;
