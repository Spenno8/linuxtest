import { useRef, useEffect, useState } from 'react';
import mapboxgl, { Map as MapboxMap } from 'mapbox-gl';
import 'mapbox-gl/dist/mapbox-gl.css';

import { Button } from '../components/ui/button';
import { Link } from "react-router-dom";

const INITIAL_CENTER: [number, number] = [144.9631608, -37.8142176]
const INITIAL_ZOOM = 10.12

type Location = {
    id: string;
    name: string;
    description: string;
    lat: number;
    lng: number;
    isSaved?: boolean;
};

type BackendPin = {
    UUID: string;
    id: string;
    pintitle: string;
    pindesc: string;
    pincolor: string;
    pinlat: string;
    pinlong: string;
};

type PinsResponse = {
    "Pin ID": BackendPin[];
};


function MapPage() {
    const mapRef = useRef<MapboxMap | null>(null);
    const mapContainerRef = useRef<HTMLDivElement | null>(null);
    //const navigate = useNavigate();

    const [center, setCenter] = useState<[number, number]>(INITIAL_CENTER);
    const [zoom, setZoom] = useState<number>(INITIAL_ZOOM);

    const [selectedLocation, setSelectedLocation] = useState<Location | null>(null);
    const markersRef = useRef<Map<string, mapboxgl.Marker>>(new Map());
    const [isAddingPin, setIsAddingPin] = useState(false);

    const [locations, setLocations] = useState<Location[]>([

    ]);

    /* ---------------- Database---------------- */
    useEffect(() => {
        const fetchPins = async () => {
            try {
                const res = await fetch('http://localhost:8080/api/UserMapPins', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        UserID: '18eea458-8222-455b-a8c2-fb0bede324c3', // replace with logged-in user UUID
                    }),
                });

                if (!res.ok) {
                    throw new Error(`HTTP ${res.status}`);
                }

                const data = await res.json();
                const typedData = data as PinsResponse;

                const mapped: Location[] = typedData["Pin ID"].map((pin) => ({

                    id: pin.id,
                    name: pin.pintitle,
                    description: pin.pindesc,
                    lat: parseFloat(pin.pinlat),
                    lng: parseFloat(pin.pinlong),
                    isSaved: true,
                }));

                setLocations(mapped);
            } catch (err) {
                console.error('Failed to fetch pins:', err);
            }
        };

        fetchPins();
    }, []);

    /*---------------------------------------*/
    const savePin = async (location: Location) => {
        const endpoint = location.isSaved
            ? 'http://localhost:8080/api/UpdateUserPin'
            : 'http://localhost:8080/api/NewUserPin';

        const res = await fetch(endpoint, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                id: location.id,
                userId: '18eea458-8222-455b-a8c2-fb0bede324c3',
                pintitle: location.name,
                pindesc: location.description,
                pincolor: 'red',
                pinlat: location.lat.toString(),
                pinlong: location.lng.toString(),
            }),
        });

        if (!res.ok) {
            throw new Error('Failed to save pin');
        }

        const data = await res.json();
        const saved = data["Pin ID"];

        setLocations((prev) =>
            prev.map((loc) =>
                loc.id === location.id
                    ? {
                        ...loc,
                        id: saved.id,   // replace temp id
                        isSaved: true,  // now persisted
                    }
                    : loc
            )
        );
    };


    /* -------------- Pop Up ---------------- */
    const createPopup = (location: Location) => {
        const container = document.createElement('div');

        container.innerHTML = `
      <h3 style="font-weight:600">${location.name}</h3>
      <p style="margin-bottom:6px">${location.description}</p>
      <button
        style="
          background:#2563eb;
          color:white;
          padding:6px 10px;
          border-radius:6px;
          font-size:12px;
          cursor:pointer;
        "
      >
        Edit Details
      </button>
    `;

        container.querySelector('button')?.addEventListener('click', () => {
            setSelectedLocation(location);
        });

        return new mapboxgl.Popup({ offset: 25 }).setDOMContent(container);
    };

    /* -------------- MAP LOAD ---------------- */
    useEffect(() => {
        mapboxgl.accessToken = import.meta.env.VITE_MAPBOX_TOKEN;

        if (!mapContainerRef.current) return;

        const map = new mapboxgl.Map({
            container: mapContainerRef.current,
            style: 'mapbox://styles/mapbox/streets-v12',
            center: INITIAL_CENTER,
            zoom: INITIAL_ZOOM,
        });

        mapRef.current = map;

        map.on('load', () => {
            map.resize();
        });

        map.on('move', () => {
            const c = map.getCenter();
            setCenter([c.lng, c.lat]);
            setZoom(map.getZoom());
        });

        return () => {
            map.remove();
        };
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    /* -------------- Save Markers ---------------- */
    useEffect(() => {
        if (!mapRef.current) return;

        const map = mapRef.current;

        // clear existing markers
        markersRef.current.forEach(marker => marker.remove());
        markersRef.current.clear();

        // add markers from state
        locations.forEach(location => {
            const marker = new mapboxgl.Marker()
                .setLngLat([location.lng, location.lat])
                .setPopup(createPopup(location))
                .addTo(map);

            markersRef.current.set(location.id, marker);
        });
    }, [locations]);

    /* ---------------- CLICK ---------------- */
    useEffect(() => {
        if (!mapRef.current) return;

        const map = mapRef.current;

        const handleClick = (e: mapboxgl.MapMouseEvent) => {
            if (!isAddingPin) return;

            const { lng, lat } = e.lngLat;

            const newLocation: Location = {
                id: crypto.randomUUID(),
                name: 'New Location',
                description: 'Click edit to update',
                lat,
                lng,
                isSaved: false,
            };

            setLocations((prev) => [...prev, newLocation]);

            setSelectedLocation(newLocation);
            setIsAddingPin(false);
        };

        map.on('click', handleClick);

        return () => {
            map.off('click', handleClick);
        };
    }, [isAddingPin]);

    /* ---------------- CURSOR ---------------- */

    useEffect(() => {
        if (!mapRef.current) return;

        mapRef.current.getCanvas().style.cursor =
            isAddingPin ? 'crosshair' : '';
    }, [isAddingPin]);

    /*--------------------------------------*/





    /* ---------------- WEBPAGE RENDERING ---------------- */
    return (
        <div className="map-page fixed inset-0 h-screen w-screen overflow-hidden">
            {/* Map Container */}
            <div
                ref={mapContainerRef}
                className="h-full w-full"
            />
            {/*Position */}
            <div className="absolute bottom-2 left-2 bg-white p-2 rounded shadow z-10 text-sm">
                Longitude: {center[0].toFixed(4)} | Latitude: {center[1].toFixed(4)} | Zoom: {zoom.toFixed(2)}
            </div>

            <div className="absolute top-2 left-2 p-2 rounded z-10 text-sm">
                <ul><li><Button><Link to="/">Home</Link></Button></li></ul>
                <Button
                    className="absolute top-12 left-2 z-10 focus:outline-2 focus:outline-offset-2 focus:outline-violet-500 active:bg-violet-700"
                    onClick={() => setIsAddingPin((prev) => !prev)}
                >{isAddingPin ? 'Cancel' : 'Add Pin'}
                </Button>
            </div>

            {/* Right detail panel */}
            {
                selectedLocation && (
                    <div className="absolute top-0 right-0 h-full w-80 bg-white shadow-xl z-20 p-4">
                        <h2 className="text-xl font-bold mb-4">Edit Location</h2>

                        <label className="block text-sm mb-1">Name</label>
                        <input
                            className="w-full border p-2 mb-3"
                            value={selectedLocation.name}
                            onChange={(e) =>
                                setSelectedLocation({
                                    ...selectedLocation,
                                    name: e.target.value,
                                })
                            }
                        />

                        <label className="block text-sm mb-1">Description</label>
                        <textarea
                            className="w-full border p-2 mb-4"
                            value={selectedLocation.description}
                            onChange={(e) =>
                                setSelectedLocation({
                                    ...selectedLocation,
                                    description: e.target.value,
                                })
                            }
                        />

                        <div className="flex gap-2">
                            <Button
                                onClick={async () => {
                                    if (!selectedLocation) return;

                                    try {
                                        await savePin(selectedLocation);
                                        setSelectedLocation(null);
                                    } catch (err) {
                                        console.error(err);
                                        alert('Failed to save pin');
                                    }
                                }}
                            >
                                Save
                            </Button>

                            <Button
                                variant="outline"
                                onClick={() => setSelectedLocation(null)}
                            >
                                Cancel
                            </Button>
                        </div>
                    </div>
                )
            }
        </div >
    );
}


export default MapPage;

