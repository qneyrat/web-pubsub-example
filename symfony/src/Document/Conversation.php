<?php
declare(strict_types=1);

namespace App\Document;

use Doctrine\ODM\MongoDB\Mapping\Annotations as MongoDB;
use Doctrine\ODM\MongoDB\PersistentCollection;
use Symfony\Component\Serializer\Annotation\Groups;

/**
 * @MongoDB\Document
 */
class Conversation
{
    /**
     * @MongoDB\Id
     * @Groups("conversation")
     */
    private $id;

    /**
     * @MongoDB\Field(type="collection")
     * @Groups("conversation")
     */
    private $users;

    /**
     * @MongoDB\EmbedMany(targetDocument="Message")
     * @Groups("conversation")
     */
    private $messages;

    /**
     * @return mixed|PersistentCollection
     */
    public function getMessages()
    {
        return $this->messages;
    }

    /**
     * @param mixed $messages
     */
    public function setMessages($messages): void
    {
        $this->messages = $messages;
    }

    /**
     * @return mixed
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * @return mixed
     */
    public function getUsers()
    {
        return $this->users;
    }

    /**
     * @param mixed $users
     */
    public function setUsers($users): void
    {
        $this->users = $users;
    }
}
