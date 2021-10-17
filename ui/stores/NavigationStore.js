import { useLocation } from "react-router-dom";

class NavigationStore {
    constructor(appStore) {
        this.appStore = appStore;
    }
    
    componentHeaders = {
        "/progress": "Progress",
        "/progress/metrics": "Metrics",
        "/timer": "Timer",
        "/workouts": "Workouts",
        "/workouts/create-workout": "Create Workout",
        "/exercises": "Exercises",
        "/profile": "Profile",
        "/profile/locations": "Locations",
        "/profile/settings": "Settings",
        "/contact": "Contact",
    };
}

export default NavigationStore;
