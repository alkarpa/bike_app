import { useState, useEffect } from "react"
import fetch_service from "../services/FetchService"
import ErrorMessage from "./ErrorMessage"
import LoadingPlaceholder from "./LoadingPlaceholder"


const RideTable = ({ stationLang }) => {
  const [rides, setRides] = useState({ loading: true })
  const [page, setPage] = useState(0)
  const [parameters, setParameters] = useState({})
  const [processing, setProcessing] = useState(true)

  useEffect(() => {
    setProcessing(true)
    const fetch_data = async () => {
      let fetched_rides = await fetch_service.getRides(parameters)
      setRides(fetched_rides)
      setProcessing(false)
    }
    fetch_data()
  }, [parameters])


  if (rides.loading) {
    return (<LoadingPlaceholder />)
  }
  if (rides.error) {
    return (
      <ErrorMessage error={rides.error} />
    )
  }


  const changePage = (p) => {
    setParameters({ ...parameters, page: p })
    setPage(p)
  }

  const changeOrder = (o) => {
    const ordering = (o === parameters.order) ? `${o}_desc` : o
    setParameters({ ...parameters, order: ordering })
  }

  const m2km = (m) => (Math.floor(m / 1000))
  const secondsToMinutes = (s) => (Math.floor(s / 60))

  const OrderableTh = ({ order, changeOrder, orderable, children, w }) => (
    <th onClick={() => changeOrder(orderable)} style={w ? { width: `${w}ch` } : {}}>
      {children}
      {order === orderable
        ? <span>&uarr;</span>
        : order === orderable + "_desc" ? <span>&darr;</span> : <>&nbsp;</>
      }
    </th>
  )

  const rides_list = rides.data.rides

  const page_size = 100
  const showing_count = rides_list.length - 1
  const first_index = page * page_size + 1
  const rides_count = rides.data.count
  const last_page = Math.floor(rides_count / page_size)

  const PageButtons = () => (
    <div style={{
      textAlign: 'center', display: 'grid',
      gridTemplateColumns: 'repeat(5, 1fr)', width: '70ch', marginLeft: '50%', transform: 'translateX(-50%)'
    }}>
      {page > 0
        ? <>
          <button onClick={() => changePage(0)}>First</button>
          <button onClick={() => changePage(page - 1)}>Prev</button>
        </>
        : <><div></div><div></div></>
      }
      <span>Page {page + 1} / {last_page + 1}</span>
      {page < last_page
        ? <>
          <button onClick={() => changePage(page + 1)}>Next</button>
          <button onClick={() => changePage(last_page)}>Last</button>
        </>
        : <><div></div><div></div></>
      }
    </div>
  )

  const DateTime = ({datestring}) => {
    const date = datestring.substring(0,10)
    const time = datestring.substring(11)
    const nowrap = {whiteSpace: 'nowrap'}
    return (
      <>
      <span style={nowrap}>{date}</span><br/>
      <span>{time}</span>
      </>
    )
  }

  const RTable = () => {

    return (
      <table style={{ fontFamily: 'monospace', width: '80ch' }}>
        <caption>Showing rides {'['}{first_index},{first_index + showing_count}{']'} of {rides_count}</caption>
        <thead style={{whiteSpace: 'nowrap'}}>
          <tr>
            <OrderableTh order={parameters['order']} w={14} changeOrder={changeOrder} orderable="departure">Departure</OrderableTh>
            <OrderableTh order={parameters['order']} w={14} changeOrder={changeOrder} orderable="return">Return</OrderableTh>
            <OrderableTh order={parameters['order']} w={22} changeOrder={changeOrder} orderable="departure_station">Departure station</OrderableTh>
            <OrderableTh order={parameters['order']} w={22} changeOrder={changeOrder} orderable="return_station">Return station</OrderableTh>
            <OrderableTh order={parameters['order']} w={10} changeOrder={changeOrder} orderable="distance">Distance</OrderableTh>
            <OrderableTh order={parameters['order']} w={10} changeOrder={changeOrder} orderable="duration">Duration</OrderableTh>
          </tr>
        </thead>
        <tbody style={{ position: 'relative', overflowX: 'hidden' }}>
          <ProcessingCover />
          {
            rides_list.map(ride => (
              <tr key={`ride_${ride.departure}_${ride.return}_${ride.departure_station_id}_${ride.return_station_id}`}>
                <td><DateTime datestring={ride.departure} /></td>
                <td><DateTime datestring={ride.return} /></td>
                <td>{stationLang[ride.departure_station_id]?.name}</td>
                <td>{stationLang[ride.return_station_id]?.name}</td>
                <td>{m2km(ride.distance)} km</td>
                <td>{secondsToMinutes(ride.duration)} min</td>
              </tr>
            ))
          }

        </tbody>
      </table>
    )
  }

  const ProcessingCover = () => {
    if (processing) {
      return (
        <tr style={{
          width: '100%', height: '100%', zIndex: '50',
          position: 'absolute', opacity: '0.5',
          backgroundImage: 'linear-gradient(gray, transparent)',
          textAlign: 'center'
        }}>
          <td colSpan={6}>
            <div style={{ height: '100%', width: '80ch' }}>
              <h1 style={{textShadow: '2px 2px white'}}>Processing</h1>

              <div className='Rotating' style={{
                width: '2em',
                height: '10em',
                backgroundColor: 'black',
                marginLeft: '50%',
              }}></div>
            </div>
          </td>
        </tr>
      )
    }
    return <></>
  }


  return (
    <div>
      <h1 style={{ textAlign: 'center' }}>Rides</h1>
      <PageButtons />
      <RTable />
    </div>

  )
}

export default RideTable