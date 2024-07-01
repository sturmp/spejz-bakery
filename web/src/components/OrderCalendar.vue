<script setup>
import { ref } from 'vue';

const schedules = ref(null);
const dayOffs = ref(null);
const datesOfWeek = ref(getDatesOfTheWeek());
const daysOfWeek = [
    'Hétfő',
    'Kedd',
    'Szerda',
    'Csütörtök',
    'Péntek',
    'Szombat',
    'Vasárnap'
];

function getDatesOfTheWeek() {
    var days = new Array();
    var current = new Date();
    var dayofWeek = current.getDay();
    if(dayofWeek == 6 || dayofWeek == 0 || (dayofWeek == 5 && current.getHours() > 12)){
        switch (dayofWeek) {
            case 5: current.setDate(current.getDate() + 3); break;
            case 6: current.setDate(current.getDate() + 2); break;
            case 0: current.setDate(current.getDate() + 1); break;
        }
    } else {
        current.setDate(current.getDate() - current.getDay() + (current.getDay() == 0 ? -6 : 1));
    }

    for(var i = 0; i < 7; i++) {
        days.push(new Date(current));
        current.setDate(current.getDate() + 1);
    }

    return days;
}

const scheduleUrl = `${import.meta.env.VITE_API_URL}/schedule`;
async function fetchBakingSchedulesAsync() {
    const requestOptions = {
        method: 'GET',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN }
    };
    schedules.value = await (await fetch(scheduleUrl, requestOptions)).json();
    schedules.value = schedules.value.filter(schedule => {
        return schedule.ReadyDate >= datesOfWeek.value[0]
            && schedule.ReadyDate <= datesOfWeek.value[6]
    })
}

const dayOffUrl = `${import.meta.env.VITE_API_URL}/dayoff`;
async function fetchDayOffsAsync() {
    const requestOptions = {
        method: 'GET',
        headers: { 'AuthToken': import.meta.env.VITE_API_AUTH_TOKEN }
    };
    dayOffs.value = await (await fetch(dayOffUrl, requestOptions)).json();
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
        && scheduleDate.getDate() == dateOfWeek.getDate();
}

function isInTheMoring(scheduleDate) {
    return scheduleDate.getHours() < 12;
}

function isInTheAfternoon(scheduleDate) {
    return scheduleDate.getHours() > 12;
}

function isDayOff(dayOfWeek) {
    const foundDayOff = dayOffs.value.find((dayOff) =>{
        return isSameDay(new Date(dayOff), datesOfWeek.value[dayOfWeek]);
    })

    return foundDayOff != undefined;
}

fetchBakingSchedulesAsync()
fetchDayOffsAsync()
</script>

<template>
    <div id="calendar" v-if="schedules != null">
        <h3>{{ datesOfWeek[0].getMonth()+1 }}.{{ datesOfWeek[0].getDate() }} - {{ datesOfWeek[6].getMonth()+1 }}.{{ datesOfWeek[6].getDate() }}</h3>
        <div class="row" v-for="(day, index) in daysOfWeek" :key="day">
            <div class="day">{{ day }}</div>
            <div class="order-rows">
                <div class="orders" v-if="anyScheduleForGivenDay(index) == undefined">
                    <div class="order">{{ isDayOff(index)? 'X' : '' }}</div>
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