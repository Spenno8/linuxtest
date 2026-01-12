import { useRef, useEffect } from 'react'
import mapboxgl, { Map } from 'mapbox-gl'

import 'mapbox-gl/dist/mapbox-gl.css';

import './map.css'

import { Button } from '../components/ui/button';
import { useNavigate } from "react-router-dom";

function App() {

    const mapRef = useRef<Map | null>(null)
    const mapContainerRef = useRef<HTMLDivElement | null>(null)
    const navigate = useNavigate()

    useEffect(() => {
        mapboxgl.accessToken = import.meta.env.VITE_MAPBOX_TOKEN
        mapRef.current = new mapboxgl.Map({
            container: mapContainerRef.current!,
            style: "mapbox://styles/mapbox/streets-v12",
            center: [144.9631608, -37.8142176],
            zoom: 10,
        });

        return () => {
            mapRef.current?.remove()
        }
    }, [])

    return (
        <>
            <div id='map-container' ref={mapContainerRef} />

            {/* Sidebar */}
            <div className="absolute top-0 left-0 h-full w-64 bg-white shadow-lg z-10 flex flex-col">
                <h2 className="text-xl font-bold p-4 border-b">Navigation</h2>
                <nav className="flex flex-col p-4 gap-2">
                    <Button onClick={() => navigate("/")}> Home</Button>
                </nav>
            </div>

        </>
    )
}

export default App
