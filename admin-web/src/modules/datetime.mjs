export function formatDate(date) {
    const month = date.getMonth()+1;
    const monthString = month < 10 ? `0${month}` : month;
    const day = date.getDate();
    const dayString = day < 10 ? `0${day}` : day;

    return `${date.getFullYear()}-${monthString}-${dayString} ${date.getHours()<=12? "Morning":"Afternoon"}`;
}

const days = [
    'Sunday',
    'Monday',
    'Tuesday',
    'Wednesday',
    'Thursday',
    'Friday',
    'Saturday'
];

export function formatDateWithDayName(date) {
    const month = date.getMonth() + 1;
    const monthString = month < 10 ? `0${month}` : month;
    const day = date.getDate();
    const dayString = day < 10 ? `0${day}` : day;

    return `${date.getFullYear()}-${monthString}-${dayString} ${days[date.getDay()]}`;
}

export function calculateStartOfWeek() {
    var today = new Date();
    today.getDay();
    return today.setDate(today.getDate()-today.getDay());
}