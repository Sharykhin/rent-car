create or replace procedure rent_car(carId uuid, consumerId uuid, dateFrom timestamptz, dateTo timestamptz)
language plpgsql
as $$
declare
foundCarId uuid;
foundConsumerId uuid;
foundRequisitionId uuid;
begin
	PERFORM pg_sleep(1);
	select c.id into foundCarId from cars c where c.id=carId for update;
	if not found then raise exception '[ERR_NOT_FOUND] Car id % was not found', carId;
    end if;

   	select c.id into foundConsumerId from consumers c where c.id=consumerId;
	if not found then raise exception '[ERR_NOT_FOUND] Consumer id % was not found', consumerId;
    end if;

    select r.id into foundRequisitionId from requisitions r where r.car_id = carId
	and (r.date_from, r.date_to ) overlaps (dateFrom, dateTo) limit 1;
    if found then raise exception '[ERR_OVERLAPPING] Provided period is overlapping';
    end if;

    insert into requisitions (car_id , consumer_id , date_from , date_to , created_at )
    values(carId, consumerId, dateFrom , dateTo, now());

   	commit;

end;$$