export const parseDate = (date: string): Date => {
    let cleanedDate = date.split(" m=")[0];
    cleanedDate = cleanedDate.replace(" UTC", "Z");
    return new Date(cleanedDate);
};