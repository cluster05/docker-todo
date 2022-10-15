import axios from "axios";


export const todoservice = axios.create({
    baseURL: `${process.env.VUE_APP_API_URL}/todo`,
});