<div id="map"></div>

<script lang="ts">
	import 'leaflet/dist/leaflet.css';
	import { onMount } from 'svelte';
	import { createEventDispatcher } from 'svelte';
	import { markers, points } from '$lib/mapMarkers';
	import type { LatLngExpression } from 'leaflet';

	const dispatch = createEventDispatcher();

	const tokyoCoords: LatLngExpression = [35.6786, 139.7637];
	
	onMount(async () => {
		// async import because import calls window
		const L = await import('leaflet');

		const map = L.map('map').setView([45.7096757,17.1851883], 4);
		L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png').addTo(map);
		map.setMaxZoom(9).setMinZoom(3);

		/*
		const marker = L.marker(markers[0].coords)
			.bindTooltip('Planning the route')
			.addTo(map);
		 marker.addEventListener('click', () => dispatch('clicked', markers[0].id));
		*/
		var polyline = L.polyline(points.map(p => p.coords), {color: 'black', opacity: 1, className: 'no-click'}).addTo(map);
		L.polyline([points[points.length-1].coords, tokyoCoords], { color: 'black', dashArray: "15 10", opacity: 0.5}).addTo(map);
		// map.fitBounds(polyline.getBounds()); // Zoom map out to see entire route

		for (const point of points) {
			const circleMarker = L.circleMarker(point.coords, {
				color: point.clickable ? 'var(--blue)' : 'black',
				fillColor: point.clickable ? 'var(--orange)' : 'white',
				fillOpacity: 1,
				radius: 7,
				className: point.clickable ? 'point' : 'no-click point'
			})
			.bindTooltip(point.hoverText)
			.addTo(map);
			if (point.clickable) {
				circleMarker.addEventListener('click', () => {
					dispatch('point-click', point.articleId);
				});
			}
		}
		L.circleMarker(tokyoCoords, {
				color: 'black',
				fillColor: 'white',
				fillOpacity: 1,
				radius: 7,
				className: 'point'
			})
			.bindTooltip('Tokyo, Japan')
			.addTo(map);
	});
</script>

<style>
	#map {
		height: calc(100dvh - var(--header-height));
	}

	:global(.no-click) {
		cursor: grab;
	}

	:global(.point) {
		outline: none;
	}
</style>