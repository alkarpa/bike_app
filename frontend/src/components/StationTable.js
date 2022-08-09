import { useState } from "react"
import StationMap from "./StationMap"
import StationView from "./StationView"
import './StationTable.css'

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

const StationFilters = ({ filters, setFilters, stations }) => {

    const stations_min_id = stations.reduce( (prev, cur) => Math.min(prev, cur.id) , stations[0].id )
    const stations_max_id = stations.reduce( (prev, cur) => Math.max(prev, cur.id) , stations[0].id )

    const [text, setText] = useState( filters.text || '')
    const [minID, setMinID] = useState(filters.minID || stations_min_id)
    const [maxID, setMaxID] = useState(filters.maxID || stations_max_id)

    const handleTextChange = (event) => {
        const value = event.target.value
        setText(value)
        setFilters({ ...filters, text: value })
    }
    const handleTextReset = () => {
        handleTextChange({target: {value: ''}})
    }

    const handleMinIDChange = (event) => {
        const value = event.target.value
        setMinID(value)
        setFilters({ ...filters, minID: value})
    }

    const handleMaxIDChange = (event) => {
        const value = event.target.value
        setMaxID(value)
        setFilters({ ...filters, maxID: value})
    }

    return (
        <div className="Panel">
            <header>Filters</header>
            <div style={{ margin: '0.5em' }}>

                <div className="Panel">
                <header>
                    Text filter
                    
                </header>
                <input value={text} onChange={handleTextChange} />
                <button className="LinkButton" onClick={handleTextReset}>Reset</button>

                </div>

                <div className="Panel">

                <header>
                    Min ID
                </header>
                <input type='number' onChange={handleMinIDChange} value={minID} min={stations_min_id} max={stations_max_id} />
                <button className="LinkButton" 
                    onClick={() => handleMinIDChange({target: {value: stations_min_id}})}>Reset</button>
                </div>
                <div className="Panel">

                <header>
                    Max ID
                   
                </header>
                <input type='number' onChange={handleMaxIDChange} value={maxID} min={stations_min_id} max={stations_max_id} />
                <button className="LinkButton" 
                    onClick={() => handleMaxIDChange({target: {value: stations_max_id}})}>Reset</button>
                </div>


            </div>
        </div>
    )
}

const StationTable = ({ lang, stations }) => {
    const [page, setPage] = useState(0)
    const [station, setStation] = useState(undefined)
    const [filters, setFilters] = useState({})
    const [orderings, setOrderings] = useState({ by: 'id', dir: 1 })

    if (station !== undefined) {
        const index = stations.findIndex(s => s.id === station.id)
        return (
            <div>
                <StationView station={station} setStation={setStation} lang={lang} stations={stations} />
                <StationMap stations={stations} filtered={stations} active_low={index} active_high={index + 1} />
            </div>
        )
    }

    const changeFilters = (f) => {
        setPage(0)
        setFilters(f)
    }

    const text = (station, key) => {
        if(station.text[lang][key]) {
            return station.text[lang][key]
        }
        return 'N/A'
    }

    let filtered = [...stations]
    if (filters.text) {
        filtered = filtered.filter((s) => JSON.stringify(s).includes(filters.text))
    }
    if (filters.minID) {
        filtered = filtered.filter((s) => s.id >= filters.minID)
    }
    if (filters.maxID) {
        filtered = filtered.filter((s) => s.id <= filters.maxID)
    }

    const sorters = {
        id: (a, b) => ((a.id - b.id) * orderings.dir),
        name: (a, b) => (a.text[lang].name.localeCompare(b.text[lang].name) * orderings.dir),
        city: (a,b) => (text(a,'city').localeCompare(text(b,'city')) * orderings.dir),
    }

    filtered.sort(sorters[orderings.by])

    const page_size = 10
    const station_low = page * page_size
    const station_high = Math.min((page + 1) * page_size, filtered.length)
    const page_slice = filtered.slice(station_low, station_high)

    const TableFiller = () => (
        <>
        {
            Array.from({length: page_size-page_slice.length}, () => 0).map( (_, i) => (
                <tr key={`table_filler_${i}`}>
                    <td colSpan={3}>&nbsp;</td>
                </tr>
            ) )
        }
        </>
    )

    const handleOrdering = (by) => {
        if (orderings.by === by) {
            setOrderings({ ...orderings, dir: orderings.dir * -1 })
        } else {
            setOrderings({ by: by, dir: 1 })
        }
    }


    return (
        <div style={{ height: '100%' }}>
            <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'space-around' }}>
                <StationFilters filters={filters} setFilters={changeFilters} stations={stations} />
                <div className="StationTable Panel">
                    <PageEnumeration page={page} setPage={setPage} page_size={page_size} content_array={filtered} />
                    <table style={{ width: '100%' }}>
                        <caption>Showing stations {'['}{station_low + 1},{station_high}{']'} out of {filtered.length}</caption>
                        <thead>
                            <tr>
                                <th style={{ width: '5em' }} onClick={() => handleOrdering('id')}>ID</th>
                                <th onClick={() => handleOrdering('name')}>Name</th>
                                <th onClick={() => handleOrdering('city')}>City</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                page_slice.map(station => (
                                    <tr key={`station_tr_${station.id}`} onClick={() => setStation(station)}>
                                        <td>{station.id}</td>
                                        <td>{text(station, 'name')}</td>
                                        <td>{text(station, 'city')}</td>
                                    </tr>
                                ))
                            }
                            {
                                <TableFiller />
                            }
                        </tbody>
                    </table>
                </div>

                <StationMap stations={stations} filtered={filtered} active_low={station_low} active_high={station_high} />
            </div>
        </div>
    )
}

export default StationTable