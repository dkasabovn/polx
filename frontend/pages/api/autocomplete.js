import axios from 'axios'

export default async function handler(req, res) {
    if (req.method === "GET") {
        const {
            query: { value }
        } = req
        try {
            const resp = await axios.get('http://localhost:6969/shills/autocomplete', { params: { value: value } })
            return res.status(200).json(resp.data)
        } catch (e) {
            return res.status(500).json({"data": null})
        }
    }
    return res.status(400).json({"data": null})
}