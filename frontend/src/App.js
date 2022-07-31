import { useEffect, useState } from 'react';
import './App.css';
import ErrorMessage from './components/ErrorMessage';
import LoadingPlaceholder from './components/LoadingPlaceholder';
import RideTable from './components/RideTable';
import StationTable from './components/StationTable';
import fetch_service from './services/FetchService';

const views = ["stations", "rides"]

const App = () => {
  const [lang, setLang] = useState('fi')
  const [stations, setStations] = useState({ loading: true })
  const [view, setView] = useState(views[0])

  useEffect(() => {
    const fetch_data = async () => {
      let stations = await fetch_service.getStations()
      setStations(stations)
    }
    fetch_data()
  }, [])

  const LanguageSelection = () => {
    const flags = {
      "se": "🇸🇪",
      "fi": "🇫🇮",
    }
    return (
      <div>
        Data language:
        {["fi", "se"].map(l => (<button key={`lb_${l}`} onClick={() => setLang(l)}>{flags[l]}</button>))}
      </div>
    )
  }

  const ViewTabs = () => (
    <div style={{ display: 'grid', gridTemplateColumns: `repeat(${views.length}, 1fr)` }}>
      {views.map(v => (<button key={`view_${v}`}
        onClick={() => setView(v)}
        style={view === v ? { backgroundColor: 'white', borderBottom: 'none' } : {}}
      >{v}</button>))}
    </div>
  )

  const View = () => {
    if (stations.loading) {
      return (<LoadingPlaceholder />)
    }
    if (stations.error) {
      return <ErrorMessage error={stations.error} />
    }

    const stations_list = stations.data
    const stationLang = stations_list.reduce((map, cur) => ({ ...map, [cur.id]: cur["text"][lang] }), {})
    return (
      <div className="View">
        {
          view === 'stations'
            ? <StationTable lang={lang} stations={stations_list} stationLang={stationLang} />
            : view === 'rides'
              ? <RideTable stationLang={stationLang} />
              : <></>
        }

      </div>
    )
  }

  return (
    <div className="App">
      <header>
        <LanguageSelection />
        <ViewTabs />
      </header>
      <View />
    </div>
  );
}

export default App;
