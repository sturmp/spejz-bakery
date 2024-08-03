<script setup>
import { ref } from 'vue';
import { fetchFromApi } from '@/modules/fetch.mjs';
import { useI18n } from 'vue-i18n';
const { t } = useI18n();

const schedules = ref(null);
const dayOffs = ref(null);
const datesOfWeek = ref(getDatesOfTheWeek());

const daysOfWeek = [
        t('monday'),
        t('tuesday'),
        t('wednesday'),
        t('thursday'),
        t('friday'),
        t('saturday'),
        t('sunday')
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
    schedules.value = await (await fetchFromApi(scheduleUrl)).json();
    schedules.value = schedules.value.filter(schedule => {
        var scheduleReadyDate = new Date(schedule.ReadyDate)
        return scheduleReadyDate >= datesOfWeek.value[0]
            && scheduleReadyDate <= datesOfWeek.value[6]
    })
}

const dayOffUrl = `${import.meta.env.VITE_API_URL}/dayoff`;
async function fetchDayOffsAsync() {
    dayOffs.value = await (await fetchFromApi(dayOffUrl)).json();
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
        return isSameDay(new Date(dayOff.Day), datesOfWeek.value[dayOfWeek]);
    })

    return foundDayOff != undefined;
}

function getFormatedDate(date) {
    return `${addPaddingZero(date.getMonth()+1)}.${addPaddingZero(date.getDate())}`;
}

function addPaddingZero(number) {
    if (number < 10) {
        return `0${number}`;
    }

    return number;
}

fetchBakingSchedulesAsync()
fetchDayOffsAsync()
</script>

<template>
    <div id="calendar" v-if="schedules != null">
        <h3>{{ `${getFormatedDate(datesOfWeek[0])} - ${getFormatedDate(datesOfWeek[6])}` }}</h3>
        <div class="row" v-for="(day, index) in daysOfWeek" :key="day">
            <div class="day">{{ day }}</div>
            <div class="order-rows">
                <div class="orders" v-if="anyScheduleForGivenDay(index) == undefined">
                    <div class="order">{{ isDayOff(index)? 'X' : '' }}</div>
                </div>
                <div class="orders" v-if="filterSchedulesForGivenDaysMorning(index).length > 0">
                    <div class="order part-of-day">{{ t('morning') }}</div>
                    <div class="order" v-for="schedule in filterSchedulesForGivenDaysMorning(index)" :key="schedule.Pastry">
                        {{ schedule.Pastry.Name }}<br>({{ schedule.Reserved }}/{{ schedule.Quantity }})
                    </div>
                </div>
                <div class="orders" v-if="filterSchedulesForGivenDaysAfternoon(index).length > 0">
                    <div class="order part-of-day">{{ t('afternoon') }}</div>
                    <div class="order" v-for="schedule in filterSchedulesForGivenDaysAfternoon(index)" :key="schedule.Pastry">
                        {{ schedule.Pastry.Name }}<br>({{ schedule.Reserved }}/{{ schedule.Quantity }})
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
    flex: 0 1 5.5em;

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
    padding: 0.2em;
}
</style>