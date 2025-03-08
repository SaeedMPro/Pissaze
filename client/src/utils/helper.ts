//import component
import moment from 'moment-jalaali';


export const jalaliMoment = (isoString: string) => {
    const date =moment(isoString);
    // Convert to Jalali and format with full details
    return date.format('jYYYY/jMMMM/jDD  [ساعت]  HH:mm'); // Return the formatted Persian date
};
