# Supermarket Checkout Kata

## The Brief

For this task we’ll build an application for a supermarket checkout. It should allow users to scan items, remove items and calculate the total price for a basket.

Each product in the system is identified using a SKU (Stock Keeping Unit) and for simplicity we’ll keep these simple by using single letters of the alphabet.

### Pricing Rules

The simplest pricing rules the software should support are:

- Price by unit, e.g. baked beans are £0.85 per can.
- Price by weight, e.g. apples are £0.65 per kg.

However there are many complex pricing variations that will potentially need consideration:

- Unit price reduction/discount, e.g. £0.20/10% off per unit
- Weight price reduction/discount, e.g. £0.20/10% off per kg
- Multi-purchase price, e.g. 3 for £10
- Multi-purchase freebie, e.g. buy 2 get 1 free
- Multi-purchase freebie (different products), e.g. buy 5 DVDs get cheapest free

Each item in the basket can have only one variation applied at a time - e.g. you can't apply 10% off and also have buy 2 get 1 free. The pricing logic should allow for variations to be prioritised.

### Product Data

Below is a table of some dummy products that we can use.

| SKU | Product     | Price    | Variations       |
| --- | ----------- | -------- | ---------------- |
| A   | Apples      | £0.65/kg |                  |
| B   | Baked Beans | £0.85    | Buy 3 get 1 free |
| C   | Cola        | £1.20    | 20% off          |
| D   | Deoderant   | £5.00    | £1.50 off        |
| E   | Eggs        | £1.20    |                  |
| F   | Fish        | £3.99    | 3 for £10        |

## Technical Requirements

This application should either be built as a UI in React or as an API in NodeJS, preferably using TypeScript. It is expected that for each **deliverable** tests will be written, using Jest or a similar framework.

We **don't** want to spend a lot of time on the project setup; we're not looking to build something that is production ready. If you have a preference for a boilerplate to use then by all means go with that. If not then keep it simple, eg for UI you can use Create React App to get a working environment with a testing framework.

```
yarn create react-app my-app --template typescript <folder_name>
```

Alternatively, you can use [Code Sandbox](https://codesandbox.io) which has plenty of templates to pick from.

## Stages

### Create A Backlog (10 minutes)

Begin by breaking the requirements down in to deliverables, these can be either user stories or technical tasks. Take some time to consider the effort required for each; feel free to estimate if that helps. The main aim is to have an idea of what can be achieved within the limited timeframe and some form of acceptance criteria that can guide the tests we write.

### Coding The Application (45 minutes)

At Birdie code quality is extremely important however, the main focus of this exercise is to see how you approach the problem and solve it collaboratively. As the time to deliver this is very limited it's imperative that we are pragmatic in our approach.

### Wash-up (5 minutes)

This is an opportunity to discuss what has been achieved and compare it to the original plan/backlog.

- What would you have done differently?
- What would you do if you had more time?
