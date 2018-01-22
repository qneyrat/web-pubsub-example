<?php
declare(strict_types=1);

namespace App\EventSubscriber\Doctrine;

use App\Entity\Message;
use App\Event\MessageEvent;
use Doctrine\Common\EventSubscriber;
use Doctrine\ORM\Event\LifecycleEventArgs;
use Symfony\Component\EventDispatcher\EventDispatcherInterface;

class MessageSubscriber implements EventSubscriber
{
    /**
     * @var EventDispatcherInterface
     */
    private $eventDispatcher;

    /**
     * @param EventDispatcherInterface $eventDispatcher
     */
    public function __construct(EventDispatcherInterface $eventDispatcher)
    {
        $this->eventDispatcher = $eventDispatcher;
    }

    /**
     * {@inheritdoc}
     */
    public function getSubscribedEvents()
    {
        return [
            'postPersist',
        ];
    }

    public function postPersist(LifecycleEventArgs $args)
    {
        if (!$args->getObject() instanceof Message) {
            return;
        }

        $this->eventDispatcher->dispatch(
            MessageEvent::CREATED,
            new MessageEvent($args->getObject())
        );
    }
}
