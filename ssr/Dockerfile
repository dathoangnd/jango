FROM node:14.16-alpine

COPY . .

RUN npm install
RUN npm run build

CMD ["npm", "run", "start"]