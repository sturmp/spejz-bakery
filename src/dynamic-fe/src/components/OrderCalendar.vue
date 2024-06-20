<script setup>
import { ref } from 'vue';

const schedules = ref(null);
const datesOfWeek = ref(getDaysOfTheWeek());
const daysOfWeek = [
    'Hétfő',
    'Kedd',
    'Szerda',
    'Csütörtök',
    'Péntek',
    'Szombat',
    'Vasárnap'
];

function getDaysOfTheWeek() {
    var days = new Array();
    var current = new Date();
    current.setDate(current.getDate() - current.getDay() + (current.getDay() == 0 ? -6 : 1));
    for(var i = 0; i < 7; i++) {
        days.push(new Date(current));
        current.setDate(current.getDate() + 1);
    }
    return days;
}

const url = "http://localhost:5555/schedule";
async function fetchBakingSchedulesAsync() {
    schedules.value = await (await fetch(url)).json();
}

function anyScheduleForGivenDay(dayOfWeek) {
    const schedule = schedules.value.find((element) => {
        return isSameDay(new Date(element.ReadyDate), datesOfWeek.value[dayOfWeek]);
    })
    return schedule;
}

function filterSchedulesForGivenDaysMorning(dayOfWeek) {
    const filteredSchedules = schedules.value.filter((schedule) => {
        const scheduleDate = new Date(schedule.ReadyDate);
        return isSameDay(scheduleDate, datesOfWeek.value[dayOfWeek])
            && isInTheMoring(scheduleDate);
    })
    return filteredSchedules;
}

function filterSchedulesForGivenDaysAfternoon(dayOfWeek) {
    const filteredSchedules = schedules.value.filter((schedule) => {
        const scheduleDate = new Date(schedule.ReadyDate);
        return isSameDay(scheduleDate, datesOfWeek.value[dayOfWeek])
            && isInTheAfternoon(scheduleDate);
    })
    return filteredSchedules;
}

function isSameDay(scheduleDate, dateOfWeek) {
    return scheduleDate.getFullYear() == dateOfWeek.getFullYear()
        && scheduleDate.getMonth() == dateOfWeek.getMonth()
        && scheduleDate.getDay() == dateOfWeek.getDay();
}

function isInTheMoring(scheduleDate) {
    return scheduleDate.getHours() < 12;
}

function isInTheAfternoon(scheduleDate) {
    return scheduleDate.getHours() > 12;
}

fetchBakingSchedulesAsync()
</script>

<template>
    <div id="calendar" v-if="schedules != null">
        <h3>{{ datesOfWeek[0].getMonth()+1 }}.{{ datesOfWeek[0].getDate() }} - {{ datesOfWeek[6].getMonth()+1 }}.{{ datesOfWeek[6].getDate() }}</h3>
        <div class="row" v-for="(day, index) in daysOfWeek" :key="day">
            <div class="day">{{ day }}</div>
            <div class="order-rows">
                <div class="orders" v-if="anyScheduleForGivenDay(index) == undefined">
                    <div class="order">X</div>
                </div>
                <div class="orders" v-if="filterSchedulesForGivenDaysMorning(index).length > 0">
                    <div class="order part-of-day">Délelőtt</div>
                    <div class="order" v-for="schedule in filterSchedulesForGivenDaysMorning(index)" :key="schedule.Pastry">
                        {{ schedule.Pastry }}<br>({{ schedule.Reserved }}/{{ schedule.Quantity }})
                    </div>
                </div>
                <div class="orders" v-if="filterSchedulesForGivenDaysAfternoon(index).length > 0">
                    <div class="order part-of-day">Délután</div>
                    <div class="order" v-for="schedule in filterSchedulesForGivenDaysAfternoon(index)" :key="schedule.Pastry">
                        {{ schedule.Pastry }}<br>({{ schedule.Reserved }}/{{ schedule.Quantity }})
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
#calendar {
    display: flex;
    flex-direction: column;
    height: 100%;
}

h3 {
    flex: 0 1 auto;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 1em;
    font-size: 1.4em;
    font-weight: 500;
}

.row {
    display: flex;
    flex: 1 1 auto;
}

#calendar .row:first-of-type {
    border-top: 0.15em dotted #2C3140;
}

.day {
    flex: 0 1 5em;

    font-weight: bold;
    border-left: 0.15em dotted #2C3140;
    border-right: 0.15em dotted #2C3140;
    border-bottom: 0.15em dotted #2C3140;

    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
}

.order-rows {
    display: flex;
    flex-direction: column;
    flex: 1 1 auto;
}

.orders {
    display: flex;
    flex: 1 1 auto;
}

.order {
    border-right: 0.15em dotted #2C3140;
    border-bottom: 0.15em dotted #2C3140;

    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;

    flex: 1 1 auto;
}

.part-of-day {
    writing-mode: vertical-rl;
    flex: 0 1 auto;
}
</style>