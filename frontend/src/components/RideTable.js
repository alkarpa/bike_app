import { useState, useEffect } from "react"

const API_path = "http://localhost:8080"

const getRides = async (parameters = {}) => {
  const ride_api = `${API_path}/ride/`
  let response = await fetch(ride_api)
  let json = await response.json()
  return json
}

const RideTable = ({stationLang}) => {
  const [rides, setRides] = useState([])

  useEffect(() => {
    const fetch_data = async () => {
      let initial_rides = await getRides()
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
                <td>{stationLang[ride.departure_station_id]?.name}</td>
                <td>{stationLang[ride.return_station_id]?.name}</td>
                <td>{m2km(ride.distance)}</td>
                <td>{secondsToMinutes(ride.duration)}</td>
              </tr>
            ))
          }

        </tbody>

      </table>

    )
}

export default RideTable