
const RideTable = ({stations, rides}) => {

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
                <td>{stations[ride.departure_station_id]}</td>
                <td>{stations[ride.return_station_id]}</td>
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