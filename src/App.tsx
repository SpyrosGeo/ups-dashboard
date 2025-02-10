import { useEffect, useState } from "react";
import "./App.css";
import { UpseResponse } from "./utils/interfaces";

function App() {
  const [isOnline,setIsOnline] = useState<boolean>(false)
  const [upsData, setUpsData] = useState<UpseResponse>();
  useEffect(() => {
    const fetchData = async () => {
      const apiResponse = await fetch("https://ups-api.thatguy.gr/api/ups");
      const data = await apiResponse?.json();
      console.log("data", data);
      setUpsData(data);
      setIsOnline(data?.STATUS ==="ONLINE")
    };
    fetchData();
  }, []);
  return (
    <>
      <div className="ups__container">
        <div className="ups__title">
        <h3>
          <span>{upsData?.MODEL}</span> UPS stats
        </h3>
    </div>
        <p className="ups__status">status: <span className={isOnline?"ups__status--online":"ups__status--offline"}> {upsData?.STATUS}</span></p>

        <p>time left: {upsData?.TIMELEFT}</p>
        <p>Charge: {upsData?.BCHARGE?.replace("Percent", "%")}</p>
        <p>Last Failure Reason: {upsData?.LASTXFER}</p>
        <p></p>
      </div>
    </>
  );
}

export default App;
