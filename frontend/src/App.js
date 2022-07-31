import { useEffect, useState } from 'react';
import './App.css';
import RideTable from './components/RideTable';
import StationTable from './components/StationTable';
import fetch_service from './services/FetchService';

const views = ["stations", "rides"]

const App = () => {
  const [lang, setLang] = useState('fi')
  const [stations, setStations] = useState([])
  const [view, setView] = useState(views[0])

  useEffect(() => {
    const fetch_data = async () => {
      let stations = await fetch_service.getStations()
      if (!stations.error) {
        setStations(stations.list)
      }
    }
    fetch_data()
  }, [])

  const flags = {
    "se": "🇸🇪",
    "fi": "🇫🇮",
  }

  const stationLang = stations.reduce((map, cur) => ({ ...map, [cur.id]: cur["text"][lang] }), {})
  return (
    <div className="App">
      <header>
        <div>
          Data language:
          {["fi", "se"].map(l => (<button key={`lb_${l}`} onClick={() => setLang(l)}>{flags[l]}</button>))}
        </div>
        <div style={{display: 'grid', gridTemplateColumns: `repeat(${views.length}, 1fr)`}}>
          {views.map(v => (<button key={`view_${v}`} 
            onClick={() => setView(v)}
            style={ view === v ? {backgroundColor: 'white', borderBottom: 'none'} : {} }
            >{v}</button>))}
        </div>
      </header>
      <div className="View">
        {
          view === 'stations'
            ? <StationTable lang={lang} stations={stations} stationLang={stationLang}/>
            : view === 'rides' 
            ? <RideTable stationLang={stationLang} /> 
            : <></>
        }

      </div>

    </div>
  );
}

export default App;
