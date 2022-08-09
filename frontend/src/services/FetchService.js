const API_path = "http://localhost:8080"

const get_fetch = async (api, param_str = '') => {
    const url = `${API_path}/${api}/${param_str}`
    const ret = { data: [] }
    try {
        const response = await fetch(url)

        if (!response.ok) {
            throw new Error(response.status)
        }

        const json = await response.json()
        ret.data = json
    } catch (error) {
        ret.error = { msg: error }
    }
    return ret
}

const getRides = async (parameters = {}) => {
    let param_str = buildRideParaStr(parameters)
    return get_fetch('ride', param_str)
}

const getStations = async () => {
    return get_fetch('station')
}

const buildRideParaStr = (parameters) => {
    const accepted_parameters = ['page', 'order', 'lang','search']
    let param = []
    for (let p of accepted_parameters) {
        if (parameters[p]) {
            param = param.concat(`${p}=${parameters[p]}`)
        }
    }

    return param.length > 0
        ? `?${param.join('&')}`
        : ''
}



const fetch_service = {
    getRides,
    getStations,
}

export default fetch_service