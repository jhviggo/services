<script lang="ts">
  import { onMount } from 'svelte';
  import { fetchUserRefuelStatistics } from '@api/statistics'

  const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul","Aug", "Sep", "Oct", "Nov", "Dec"];
  const weeks = Array.from({length: 53}, (_, i) => i + 1)
  let refuelStatistics = {};
  let canvasMonthlyCost;
  let canvasWeeklyCost;
  let canvasMonthlyTrip;
  let canvasWeeklyTrip;
  let statYear;

  const valueOptionCurrency = {
    scales: {
      y: {
        ticks: {
          callback: function(value, index, ticks) {
            return `${value}kr`;
          }
        }
      }
    }
  };

  const valueOptionDistance = {
    scales: {
      y: {
        ticks: {
          callback: function(value, index, ticks) {
            return `${value}km`;
          }
        }
      }
    }
  };

  onMount(async () => {
    const year = (new Date()).getFullYear();
    refuelStatistics = await fetchUserRefuelStatistics(year);
    statYear = refuelStatistics.year;
    console.log(year, refuelStatistics);

    // monthly cost
    const monthlyCostSums = new Array(12);
    refuelStatistics.monthly?.forEach(m => {
      monthlyCostSums[m.date_value - 1] = m.sum_cost
    });
    updateChart(canvasMonthlyCost, months, "Cost in DKK", monthlyCostSums, valueOptionCurrency);

    // weekly cost
    const weeklyCostSums = new Array(53);
    refuelStatistics.weekly?.forEach(w => {
      weeklyCostSums[w.date_value - 1] = w.sum_cost
    });

    updateChart(canvasWeeklyCost, weeks, "Cost in DKK", weeklyCostSums, valueOptionCurrency);

    // monthly trip
    const monthlyTripSums = new Array(12);
    refuelStatistics.monthly?.forEach(m => {
      monthlyTripSums[m.date_value - 1] = m.sum_trip
    });
    updateChart(canvasMonthlyTrip, months, "Trip distance in KM", monthlyTripSums, valueOptionDistance);

    // weekly trip
    const weeklyTripSums = new Array(53);
    refuelStatistics.weekly?.forEach(w => {
      weeklyTripSums[w.date_value - 1] = w.sum_trip
    });

    updateChart(canvasWeeklyTrip, weeks, "Trip distance in KM", weeklyTripSums, valueOptionDistance);
  });


  function updateChart(chartCanvas, labels, label, data, options) {
    new window.Chart(chartCanvas, {
      type: 'bar',
      data: {
        labels: labels,
        datasets: [{
          label: label,
          data: data,
          borderWidth: 1
        }]
      },
      options: options
    });
  }
</script>

<svelte:head>
  <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
</svelte:head>

<div>
  <div class="stat">
    <div class="stat-header">Year:</div>
    <div class="stat-value">{ refuelStatistics.year }</div>

    <div class="stat-header">Total cost:</div>
    <div class="stat-value">{ refuelStatistics.total_cost }kr</div>

    <div class="stat-header">Total distance:</div>
    <div class="stat-value">{ refuelStatistics.total_trip }km</div>

    <div class="stat-header">Refuels:</div>
    <div class="stat-value">{ refuelStatistics.total_refuels }</div>
  </div>

  <h2>Monthly stats</h2>
  <h3>Cost</h3>
  <canvas class="graph" bind:this={canvasMonthlyCost}></canvas>

  <h3>Trip distance</h3>
  <canvas class="graph" bind:this={canvasMonthlyTrip}></canvas>

  
  <h2>Weekly stats</h2>
  <h3>Cost</h3>
  <canvas class="graph" bind:this={canvasWeeklyCost}></canvas>

  <h3>Trip distance</h3>
  <canvas class="graph" bind:this={canvasWeeklyTrip}></canvas>
</div>

<style>
  .graph {
    max-height: 25rem;
    max-width: 100%;
  }

  .stat {
    display: grid;
    grid-template-columns: repeat(2, auto);
  }

  .stat-header {
    color: var(--primary-blue);
    font-weight: bold;
  }
</style>
