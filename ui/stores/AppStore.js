import NavigationStore from "./NavigationStore";

class AppStore {
    constructor() {
        this.navigationStore = new NavigationStore(this);
    }
};

export default AppStore;
