import { useState, useEffect } from "react"
import fetch_service from "../services/FetchService"


const RideTable = ({ stationLang }) => {
  const [rides, setRides] = useState([])
  const [page, setPage] = useState(0)
  const [parameters, setParameters] = useState({})

  useEffect(() => {
    const fetch_data = async () => {
      let fetched_rides = await fetch_service.getRides(parameters)
      if ( !fetched_rides.error ) {
        setRides(fetched_rides.list)
      }
    }
    fetch_data()
  }, [parameters])


  if (!rides || rides.length < 1) {
    return (
      <div>
        Nothing to show
      </div>
    )
  }


  const changePage = (p) => {
    setParameters({ ...parameters, page: p })
    setPage(p)
  }

  const changeOrder = (o) => {
    const ordering = (o === parameters.order) ? `${o}_desc` : o
    setParameters({...parameters, order: ordering})
  }

  const m2km = (m) => (Math.floor(m / 1000))
  const secondsToMinutes = (s) => (Math.floor(s / 60))

  const OrderableTh = ({order, changeOrder, orderable, children}) => (
    <th onClick={() => changeOrder(orderable)}>
      {children}
      { order === orderable
      ? <span>&uarr;</span>
      : order === orderable+"_desc" ? <span>&darr;</span> : <></>
     }
    </th>
  )

  return (
    <div>

      <div>
        {page > 0
          ? <>
          <button onClick={() => changePage(0)}>First</button>
          <button onClick={() => changePage(page - 1)}>Prev</button>
          </>
          : <></>
        }
        <span>Page {page}</span>
        <button onClick={() => changePage(page + 1)}>Next</button>
      </div>
      <table style={{ fontFamily: 'monospace' }}>
        <thead>
          <tr>
            <OrderableTh order={parameters['order']} changeOrder={changeOrder} orderable="departure">Departure</OrderableTh>
            <OrderableTh order={parameters['order']} changeOrder={changeOrder} orderable="return">Return</OrderableTh>
            <OrderableTh order={parameters['order']} changeOrder={changeOrder} orderable="departure_station">Departure station</OrderableTh>
            <OrderableTh order={parameters['order']} changeOrder={changeOrder} orderable="return_station">Return station</OrderableTh>
            <OrderableTh order={parameters['order']} changeOrder={changeOrder} orderable="distance">Distance</OrderableTh>
            <OrderableTh order={parameters['order']} changeOrder={changeOrder} orderable="duration">Duration</OrderableTh>
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
    </div>

  )
}

export default RideTable