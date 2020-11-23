
/**
 * This file is automatically generated using launchdarkly-code-generation tool.
 * Do not modify this file -- YOUR CHANGES WILL BE ERASED!
 */

import { LDUser, LDClient } from 'launchdarkly-node-server-sdk';

/** Class containing functions of all LaunchDarkly flags with default values mapped. */
export class LDFlags {
  private readonly ldClient: LDClient;
  /**
   * @param {LDClient} ldClient - pass in a LDClient for all flags to use in evaluation.
   */
  constructor(ldClient: LDClient) {
    this.ldClient = ldClient;
  }  /**
   * Show: Live Tiles
   *
   * Recommendation service live tiles for primary homepage placement 
   */
  async showWidgets(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('show-widgets', user, );
  }

  /**
   * Release: Dark Theme
   *
   * Newly designed Dark Theme option for users.
   */
  async darkTheme(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('dark-theme', user, );
  }

  /**
   * Enable: Chatbox
   *
   * Enable chatbox for clients
   */
  async chatbox(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('chatbox', user, );
  }

  /**
   * Beta UI
   *
   * Beta UI Features
   */
  async betaUi(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('beta-ui', user, );
  }

  /**
   * Show: NPS Survey
   *
   * This feature flag controls that layout and experience of how the NPS score will be delivered to end users.
   */
  async showNpsSurvery(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('show-nps-survery', user, );
  }

  /**
   * Show: Block
   *
   * Adds a Glassy block to your Chrome Extension enabled sites
   */
  async showBlock(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('show-block', user, );
  }

  /**
   * Configure: Registration Button Color
   *
   * This flag controls the display color of the registration button on the main landing page.
   */
  async registrationButtonColor(user: LDUser): Promise<string> {
    return this.ldClient.variation('registration-button-color', user, '#28A745');
  }

  /**
   * Execute: Force Page Reload
   *
   * Force page reload as last resort
   */
  async forcePageReload(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('force-page-reload', user, );
  }

  /**
   * React Manifest
   *
   */
  async reactManefest(user: LDUser): Promise<object> {
    return this.ldClient.variation('react-manefest', user, [{"age":25,"description":"Singer Actress","id":5,"image":"https://images.unsplash.com/photo-1564819337735-bd75ca98c507?ixlib=rb-1.2.1\u0026q=85\u0026fm=jpg\u0026crop=entropy\u0026cs=srgb\u0026ixid=eyJhcHBfaWQiOjE0NTg5fQ","name":"Miley Cyrus"},{"age":43,"description":"Author. Writer.","id":6,"image":"https://images.unsplash.com/photo-1564819337735-bd75ca98c507?ixlib=rb-1.2.1\u0026q=85\u0026fm=jpg\u0026crop=entropy\u0026cs=srgb\u0026ixid=eyJhcHBfaWQiOjE0NTg5fQ","name":"Oscar Wilde"},{"age":90,"description":"Hobbit. Thief","id":7,"image":"https://www.naschools.net/cms/lib/MA02202086/Centricity/Domain/1358/mickey-mouse.jpg","name":"Bilbo Bagins"},{"age":80,"description":"Drummer. Barman. Entrepreneur.","id":8,"image":"https://media.entertainmentearth.com/assets/images/4cd14c657232481586e4c930497a0519lg.jpg","name":"George Leaver"}]);
  }

  /**
   * The sites on fire
   *
   */
  async theSitesOnFire(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('the-sites-on-fire', user, );
  }

  /**
   * DM: Product Order
   *
   */
  async dmProductOrder(user: LDUser): Promise<object> {
    return this.ldClient.variation('dm-product-order', user, [{"id":0,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/2001-PIZ-Catalan-Chicken.jpg?v=d3cbe105139dfd554931b1984282e6dd","name":"Catalan Chicken \u0026 Chorizo","price":"19.99","size":"Large 13.5\" £19.99"},{"id":1,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Cheeseburger-Pizza.jpg?v=99d802d42befe5d1de79eb4ae37a3efc","name":"The Cheeseburger","price":"19.99","size":"Large 13.5\" £19.99"},{"id":2,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Meatfielder.jpg?v=79d95ba85e33101feeeb631f64e7fdf7","name":"The Meatfielder","price":"19.99","size":"Large 13.5\" £19.99"},{"id":3,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Cheese-and-Tomato.jpg?v=e7295a1e0db941c6b899efa2bd996bce","name":"Original Cheese \u0026 Tomato","price":"15.99","size":"Large 13.5\" £15.99"},{"id":4,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Mighty-Meaty.jpg?v=069c19724f201aa506753329c3af9a2b","name":"Mighty Meaty®","price":"19.99","size":"Large 13.5\" £19.99"},{"id":5,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Pepperoni-Passion.jpg?v=6194a868836905b257d5d4173db0889a","name":"Pepperoni Passion®","price":"19.99","size":"Large 13.5\" £19.99"},{"id":6,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Texas-BBQ.jpg?v=0c1787dd4a3ad00e87a09afef6d2bab6","name":"Texas BBQ®","price":"19.99","size":"Large 13.5\" £19.99"},{"id":7,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Vegi-Supreme.jpg?v=319295e6044a484084eb9dc994368f8e","name":"Vegi Supreme","price":"19.99","size":"Large 13.5\" £19.99"},{"id":8,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-American-Hot.jpg?v=fc55e97aed260db1febe052a0f4c26ab","name":"American Hot","price":"18.99","size":"Large 13.5\" £18.99"},{"id":9,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Chicken-Feast.jpg?v=92b663ff168d3c535d91e50a2bd461c9","name":"Chicken Feast","price":"18.99","size":"Large 13.5\" £18.99"},{"id":10,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Tandoori-Hot.jpg?v=524490372ba23ffb37ab9be0571013e9","name":"Domino’s Tandoori Hot®","price":"19.99","size":"Large 13.5\" £19.99"},{"id":11,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Farmhouse.jpg?v=0fb32d97e8102570507f4b20e9529815","name":"Farmhouse","price":"18.99","size":"Large 13.5\" £18.99"},{"id":12,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Hawaiian.jpg?v=77814d314a14f2c10b486e8c52996388","name":"Hawaiian","price":"18.99","size":"Large 13.5\" £18.99"},{"id":13,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Hot-and-Spicy.jpg?v=94699db5c550982df400eaf8fb15a1bb","name":"Hot \u0026 Spicy","price":"19.99","size":"Large 13.5\" £19.99"},{"id":14,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Meateor.jpg?v=a0dbb8bd83ca16d528b991f7653c771a","name":"Meateor™","price":"19.99","size":"Large 13.5\" £19.99"},{"id":15,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Ranch-BBQ.jpg?v=42aa13d345423defdae3e2061a5f536b","name":"Ranch BBQ","price":"19.99","size":"Large 13.5\" £19.99"},{"id":16,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Sizzler.jpg?v=cab12d9ab7f874fc5d8b69d7a613f6d9","name":"The Sizzler","price":"19.99","size":"Large 13.5\" £19.99"},{"id":17,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Tuna-Supreme.jpg?v=b0fab717312688ee13610c44b247de48","name":"Tuna Supreme","price":"18.99","size":"Large 13.5\" £18.99"},{"id":18,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Veg-A-Roma.jpg?v=1151453f7fddb4008e9b1b1339aa2f17","name":"Veg-a-Roma","price":"19.99","size":"Large 13.5\" £19.99"},{"id":19,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Vegi-Volcano.jpg?v=02483f82c204351fa18476259ffb1909","name":"Vegi Volcano","price":"19.99","size":"Large 13.5\" £19.99"},{"id":20,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Chicken-Delight-Pizza.jpg?v=50f46823eba41b55d78f0ce941334dec","name":"Delight Chicken","price":"19.99","size":"Large 13.5\" £19.99"},{"id":21,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Vegi-Delight.jpg?v=dde9844c03347b264c0b1758d996d740","name":"Delight Vegi","price":"19.99","size":"Large 13.5\" £19.99"},{"id":22,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-CYO-Delight.jpg?v=77d4999078c8651bac800fe80edd6baa","name":"Delight Create Your Own","price":"14.99","size":"Large 13.5\" £14.99"}]);
  }

  /**
   * Swimlanes
   *
   */
  async swimlanes(user: LDUser): Promise<number> {
    return this.ldClient.variation('swimlanes', user, 1);
  }

  /**
   * Show: Deal Banner
   *
   */
  async dealBanner(user: LDUser): Promise<string> {
    return this.ldClient.variation('deal-banner', user, 'Premium');
  }

  /**
   * Enable: Live Tiles Backend
   *
   */
  async liveTilesBackend(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('live-tiles-backend', user, );
  }

  /**
   * Show: Covid Banner
   *
   * Show/Hide Covid Banner
   */
  async covidBanner(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('covid-banner', user, );
  }

  /**
   * Control-DB-Type
   *
   */
  async controlDbType(user: LDUser): Promise<string> {
    return this.ldClient.variation('control-db-type', user, 'PrivateDB');
  }

  /**
   * Sort Order Demo
   *
   */
  async sortOrderDemo(user: LDUser): Promise<object> {
    return this.ldClient.variation('sort-order-demo', user, [{"id":0,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/2001-PIZ-Catalan-Chicken.jpg?v=d3cbe105139dfd554931b1984282e6dd","name":"Catalan Chicken \u0026 Chorizo","price":"19.99","size":"Large 13.5\" £19.99"},{"id":1,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Cheeseburger-Pizza.jpg?v=99d802d42befe5d1de79eb4ae37a3efc","name":"The Cheeseburger","price":"19.99","size":"Large 13.5\" £19.99"},{"id":2,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Meatfielder.jpg?v=79d95ba85e33101feeeb631f64e7fdf7","name":"The Meatfielder","price":"19.99","size":"Large 13.5\" £19.99"},{"id":3,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Cheese-and-Tomato.jpg?v=e7295a1e0db941c6b899efa2bd996bce","name":"Original Cheese \u0026 Tomato","price":"15.99","size":"Large 13.5\" £15.99"},{"id":4,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Mighty-Meaty.jpg?v=069c19724f201aa506753329c3af9a2b","name":"Mighty Meaty®","price":"19.99","size":"Large 13.5\" £19.99"},{"id":5,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Pepperoni-Passion.jpg?v=6194a868836905b257d5d4173db0889a","name":"Pepperoni Passion®","price":"19.99","size":"Large 13.5\" £19.99"},{"id":6,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Texas-BBQ.jpg?v=0c1787dd4a3ad00e87a09afef6d2bab6","name":"Texas BBQ®","price":"19.99","size":"Large 13.5\" £19.99"},{"id":7,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Vegi-Supreme.jpg?v=319295e6044a484084eb9dc994368f8e","name":"Vegi Supreme","price":"19.99","size":"Large 13.5\" £19.99"},{"id":8,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-American-Hot.jpg?v=fc55e97aed260db1febe052a0f4c26ab","name":"American Hot","price":"18.99","size":"Large 13.5\" £18.99"},{"id":9,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Chicken-Feast.jpg?v=92b663ff168d3c535d91e50a2bd461c9","name":"Chicken Feast","price":"18.99","size":"Large 13.5\" £18.99"},{"id":10,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Tandoori-Hot.jpg?v=524490372ba23ffb37ab9be0571013e9","name":"Domino’s Tandoori Hot®","price":"19.99","size":"Large 13.5\" £19.99"},{"id":11,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Farmhouse.jpg?v=0fb32d97e8102570507f4b20e9529815","name":"Farmhouse","price":"18.99","size":"Large 13.5\" £18.99"},{"id":12,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Hawaiian.jpg?v=77814d314a14f2c10b486e8c52996388","name":"Hawaiian","price":"18.99","size":"Large 13.5\" £18.99"},{"id":13,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Hot-and-Spicy.jpg?v=94699db5c550982df400eaf8fb15a1bb","name":"Hot \u0026 Spicy","price":"19.99","size":"Large 13.5\" £19.99"},{"id":14,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Meateor.jpg?v=a0dbb8bd83ca16d528b991f7653c771a","name":"Meateor™","price":"19.99","size":"Large 13.5\" £19.99"},{"id":15,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Ranch-BBQ.jpg?v=42aa13d345423defdae3e2061a5f536b","name":"Ranch BBQ","price":"19.99","size":"Large 13.5\" £19.99"},{"id":16,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Sizzler.jpg?v=cab12d9ab7f874fc5d8b69d7a613f6d9","name":"The Sizzler","price":"19.99","size":"Large 13.5\" £19.99"},{"id":17,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Tuna-Supreme.jpg?v=b0fab717312688ee13610c44b247de48","name":"Tuna Supreme","price":"18.99","size":"Large 13.5\" £18.99"},{"id":18,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Veg-A-Roma.jpg?v=1151453f7fddb4008e9b1b1339aa2f17","name":"Veg-a-Roma","price":"19.99","size":"Large 13.5\" £19.99"},{"id":19,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Vegi-Volcano.jpg?v=02483f82c204351fa18476259ffb1909","name":"Vegi Volcano","price":"19.99","size":"Large 13.5\" £19.99"},{"id":20,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Chicken-Delight-Pizza.jpg?v=50f46823eba41b55d78f0ce941334dec","name":"Delight Chicken","price":"19.99","size":"Large 13.5\" £19.99"},{"id":21,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-Vegi-Delight.jpg?v=dde9844c03347b264c0b1758d996d740","name":"Delight Vegi","price":"19.99","size":"Large 13.5\" £19.99"},{"id":22,"image":"https://www.dominos.co.uk/Content/images/Products/GB/Pizza/256x256/1907-PIZ-CYO-Delight.jpg?v=77d4999078c8651bac800fe80edd6baa","name":"Delight Create Your Own","price":"14.99","size":"Large 13.5\" £14.99"}]);
  }

  /**
   * search bar
   *
   */
  async searchBar(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('search-bar', user, );
  }

  /**
   * Delivery Options
   *
   * This is used to rollout and control different delivery options for customers
   */
  async deliveryOptions(user: LDUser): Promise<string> {
    return this.ldClient.variation('delivery-options', user, 'pickup');
  }

  /**
   * New Algorithm
   *
   * New search algorithm
   */
  async newAlgorithm(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('new-algorithm', user, );
  }
//LOCAL_LCG_FLAGS_BEGIN
  async myTestFlag(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('my-test-flag', user, true);
  }

  async anotherNewFlag(user: LDUser): Promise<boolean> {
    return this.ldClient.variation('another-new-flag', user, false);
  }

//LOCAL_LCG_FLAGS_END
}