import axios from "axios";

// export axios client, so it can be imported from the components
const reqClient = axios.create({
  baseURL: "http://localhost:8000/api"
});

export default reqClient;
