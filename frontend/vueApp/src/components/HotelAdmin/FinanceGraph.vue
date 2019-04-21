<script>
    import { Bar, mixins } from 'vue-chartjs'
    const { reactiveProp } = mixins;

    export default {
        extends: Bar,
        mixins: [reactiveProp],
        props: ["item"],
        data: function () {
            return {
                options: {
                    scales: {
                        yAxes: [{
                            ticks: {
                                beginAtZero: true
                            },
                            gridLines: {
                                display: true
                            }
                        }],
                        xAxes: [{
                            ticks: {
                                beginAtZero: true
                            },
                            gridLines: {
                                display: false
                            }
                        }]
                    },
                    legend: {
                        display: true
                    },
                    tooltips: {
                        enabled: true,
                        mode: 'single',
                        callbacks: {
                            label: function(tooltipItems) {
                                if (this["_data"]["field"] === "visitors" && tooltipItems.yLabel === 1) {
                                    return tooltipItems.yLabel + " visitor";
                                }
                                return tooltipItems.yLabel + " " + this["_data"]["field"];
                            }
                        }
                    },
                    responsive: true,
                    maintainAspectRatio: false,
                    height: 200,
                },
            }
        },
        mounted () {
            this.renderChart(this.chartData, this.options)
        }
    }
</script>

<style scoped>

</style>