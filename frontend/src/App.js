import { useEffect, useState } from 'react';
import './App.css';

const API_path = "http://localhost:8080"

const getStations = async () => {
  const station_api = `${API_path}/station/`
  let response = await fetch(station_api)
  let json = await response.json()

  const stations = json.reduce( (map, cur) => ({...map, [cur.id]: cur.name}), {})
  return stations
}

const getRides = async (parameters = {}) => {
  const ride_api = `${API_path}/ride/`
  let response = await fetch(ride_api)
  let json = await response.json()
  return json
}

function App() {
  const [stations, setStations] = useState({})
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

  if (!rides || rides.length < 1) {
    return (
      <div>
        Nothing to show
      </div>
    )
  }

  const m2km = (m) => (Math.floor(m/1000))
  const secondsToMinutes = (s) => ( Math.floor(s/60) )

  return (
    <div className="App">
      <table style={{fontFamily: 'monospace'}}>
        <thead>
          <tr>
            <th>Departure</th>
            <th>Return</th>
            <th>Departure station</th>
            <th>Return station</th>
            <th>Distance (km)</th>
            <th>Duration (min)</th>
          </tr>
        </thead>
        <tbody>
          {
            rides.map(ride => (
              <tr key={`ride_${ride.departure}_${ride.return}`}>
                <td>{ride.departure}</td>
                <td>{ride.return}</td>
                <td>{stations[ride.departure_station_id]}</td>
                <td>{stations[ride.return_station_id]}</td>
                <td>{m2km(ride.distance)}</td>
                <td>{secondsToMinutes(ride.duration)}</td>
              </tr>
            ))
          }

        </tbody>

      </table>


    </div>
  );
}

export default App;
