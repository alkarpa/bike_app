import { useState } from "react"
import StationMap from "./StationMap"
import StationView from "./StationView"

const PageButton = ({ activePage, setActivePage, buttonPage }) => {

    if (buttonPage === '...') {
        return (
            <span style={{ textAlign: 'center' }}>
                ...
            </span>
        )
    }

    return (
        <button
            onClick={() => setActivePage(buttonPage)}
            style={buttonPage === activePage ? { backgroundColor: 'white' } : {}}
        >{buttonPage + 1}</button>
    )
}

const PageEnumeration = ({ page, setPage, page_size, content_array }) => {

    const page_count = Math.ceil(content_array.length / page_size)
    const page_enumeration = Array.from({ length: page_count }, (_, i) => i)
    const page_links_length = 9


    const last_page = page_count - 1
    const distance_from_page = (a, b) => (Math.abs(page - a) - Math.abs(page - b))
    let sorty = [...page_enumeration].sort(distance_from_page)
        .slice(0, page_links_length).sort((a, b) => a - b);

    const edge_values = (array, index, value) => {
        return array.map((a, i) => {
            if (i === index) return value
            if (Math.abs(index - i) === 1 && Math.abs(value - a) > 1) return '...'
            return a
        })
    }
    sorty = edge_values(sorty, 0, 0)
    sorty = edge_values(sorty, sorty.length - 1, last_page)

    return (
        <div style={{ display: 'grid', gridTemplateColumns: `repeat(${page_links_length}, 1fr)` }}>
            {
                sorty.map((p, i) => (
                    <PageButton key={`station_table_page_${p}_${i}`} activePage={page} setActivePage={setPage} buttonPage={p} />
                ))
            }


        </div>
    )
}

const StationTable = ({ lang, stations }) => {
    const [page, setPage] = useState(0)
    const [station, setStation] = useState(undefined)

    if (station !== undefined) {
        const index = stations.findIndex( s => s.id === station.id )
        return (
            <div>
                <StationView station={station} setStation={setStation} lang={lang}/>
                <StationMap stations={stations} active_low={index} active_high={index+1} />
            </div>
        )
    }

    const page_size = 10
    const station_low = page * page_size
    const station_high = Math.min((page + 1) * page_size, stations.length)
    const page_slice = stations.slice(station_low, station_high)

    

    return (
        <div>
            <h2>Stations</h2>
            <PageEnumeration page={page} setPage={setPage} page_size={page_size} content_array={stations} />
            <table style={{ minWidth: '20em',maxWidth: '70em' }}>
                <caption>Showing stations from {station_low + 1} to {station_high} out of {stations.length}</caption>
                <thead>
                    <tr>
                        <th style={{width: '5em'}}>ID</th>
                        <th>Name</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        page_slice.map(station => (
                            <tr key={`station_tr_${station.id}`} onClick={() => setStation(station)}>
                                <td>{station.id}</td>
                                <td>{station.text[lang].name}</td>
                            </tr>
                        ))
                    }
                </tbody>
            </table>
            <StationMap stations={stations} active_low={station_low} active_high={station_high} />
        </div>
    )
}

export default StationTable