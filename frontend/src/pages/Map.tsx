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
        {
            id: '1',
            name: 'Melbourne CBD',
            description: 'Central business district',
            lat: -37.8142176,
            lng: 144.9631608,
        },
        {
            id: '2',
            name: 'South Location',
            description: 'South of the city',
            lat: -38.8142176,
            lng: 144.9631608,
        },
    ]);


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

            // Initial markers
            locations.forEach((location) => {
                const marker = new mapboxgl.Marker()
                    .setLngLat([location.lng, location.lat])
                    .setPopup(createPopup(location))
                    .addTo(map);

                markersRef.current.set(location.id, marker);
            });
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

    /* ---------------- CLICK ---------------- */
    useEffect(() => {
        if (!mapRef.current) return;

        const map = mapRef.current;

        const handleClick = (e: mapboxgl.MapMouseEvent) => {
            if (!isAddingPin) {
                setSelectedLocation(null);
                return;
            }

            const { lng, lat } = e.lngLat;

            const newLocation: Location = {
                id: crypto.randomUUID(),
                name: 'New Location',
                description: 'Click edit to update',
                lat,
                lng,
            };

            setLocations((prev) => [...prev, newLocation]);

            const marker = new mapboxgl.Marker()
                .setLngLat([lng, lat])
                .setPopup(createPopup(newLocation))
                .addTo(map);

            markersRef.current.set(newLocation.id, marker);

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
                                onClick={() => {
                                    setLocations((prev) =>
                                        prev.map((loc) =>
                                            loc.id === selectedLocation.id
                                                ? { ...selectedLocation } // ensure new object
                                                : loc
                                        )
                                    );

                                    const marker = markersRef.current.get(selectedLocation.id);
                                    if (marker) {
                                        marker.setPopup(createPopup({ ...selectedLocation }));
                                        marker.getPopup()?.remove(); // close popup
                                    }

                                    setSelectedLocation(null);
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

