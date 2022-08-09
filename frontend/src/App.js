import { useEffect, useState } from 'react';
import './App.css';
import ErrorMessage from './components/ErrorMessage';
import LoadingPlaceholder from './components/LoadingPlaceholder';
import RideTable from './components/RideTable';
import StationTable from './components/StationTable';
import fetch_service from './services/FetchService';

const views = [{id:"stations", title: 'Stations'}, {id:"rides", title:'Bike rides'}]

const App = () => {
  const [lang, setLang] = useState('fi')
  const [stations, setStations] = useState({ loading: true })
  const [view, setView] = useState(views[0].id)

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
      <div style={{textAlign: 'center'}}>
        <span style={{whiteSpace: 'nowrap'}}>Data language:</span>
        {["fi", "se"].map(l => (
        <button key={`lb_${l}`} 
                onClick={() => setLang(l)}
                style={ l === lang ? {backgroundColor: 'white'} : {} }
        >{flags[l]}
        </button>
        ))}
      </div>
    )
  }

  const ViewTabs = () => (
    <div style={{ display: 'grid', gridTemplateColumns: `repeat(${views.length}, 1fr)` }}>
      {views.map(v => (<button key={`view_${v.id}`}
        onClick={() => setView(v.id)}
        className={ view === v.id ? 'Tab ActiveTab' : 'Tab' }
        /*style={view === v ? { backgroundColor: 'white', borderBottom: 'none' } : {}}*/
      >{v.title}</button>))}
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
    switch(view) {
      case 'stations': return <StationTable lang={lang} stations={stations_list} />
      case 'rides': return <RideTable lang={lang}/>
      default: <ErrorMessage error={"Unknown view"} />
    }
  }

  return (
    <div className="App">
      <header className="AppHeader">
        <LanguageSelection />
        <ViewTabs />
      </header>
      <div className='View'>
        <View />
      </div>
      
    </div>
  );
}

export default App;
