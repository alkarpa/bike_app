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

const StationFilters = ({filters, setFilters}) => {

    const [text, setText] = useState('')

    const handleTextChange = (event) => {
        const value = event.target.value
        setText(value)
        setFilters( {...filters, text: value} )
    }

    return (
        <div>
            <h4>Filters</h4>
            <label>
            Text filter
            <input value={text} onChange={handleTextChange} />
            </label>
        </div>
    )
}

const StationTable = ({ lang, stations, stationLang }) => {
    const [page, setPage] = useState(0)
    const [station, setStation] = useState(undefined)
    const [filters, setFilters] = useState({})
    const [orderings, setOrderings] = useState({by: 'id',dir:1})

    if (station !== undefined) {
        const index = stations.findIndex( s => s.id === station.id )
        return (
            <div>
                <StationView station={station} setStation={setStation} lang={lang} stationLang={stationLang}/>
                <StationMap stations={stations} active_low={index} active_high={index+1} />
            </div>
        )
    }

    const changeFilters = (f) => {
        setPage(0)
        setFilters(f)
    }

    let filtered = [...stations]
    if ( filters.text ) {
        filtered = filtered.filter( (s) => JSON.stringify(s).includes(filters.text) )
    }

    const sorters = {
        id: (a,b) => ( (a.id-b.id)*orderings.dir ),
        name: (a,b) => (a.text[lang].name.localeCompare(b.text[lang].name)*orderings.dir)
    }

    filtered.sort( sorters[orderings.by] )

    const page_size = 10
    const station_low = page * page_size
    const station_high = Math.min((page + 1) * page_size, filtered.length)
    const page_slice = filtered.slice(station_low, station_high)

    const handleOrdering = (by) => {
        if (orderings.by === by) {
            setOrderings( {...orderings, dir: orderings.dir * -1} )
        } else {
            setOrderings( { by: by, dir: 1 } )
        }
    }
    

    return (
        <div>
            <h2>Stations</h2>
            <StationFilters filters={filters} setFilters={changeFilters} />
            <PageEnumeration page={page} setPage={setPage} page_size={page_size} content_array={filtered} />
            <table style={{ minWidth: '20em',maxWidth: '70em' }}>
                <caption>Showing stations {'['}{station_low + 1},{station_high}{']'} out of {filtered.length}</caption>
                <thead>
                    <tr>
                        <th style={{width: '5em'}} onClick={() => handleOrdering('id')}>ID</th>
                        <th onClick={() => handleOrdering('name')}>Name</th>
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
            <StationMap stations={filtered} active_low={station_low} active_high={station_high} />
        </div>
    )
}

export default StationTable